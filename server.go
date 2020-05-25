package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"path"
	"strings"
	"time"

	"gocloud.dev/blob"
	"gocloud.dev/gcerrors"
	"google.golang.org/grpc"
)

type server struct {
	httpServer *http.Server
	grpcServer *grpc.Server

	bucket *blob.Bucket
}

type handler func(w http.ResponseWriter, r *http.Request) error

func (s *server) handleErr(f handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			log.Println(err) // TODO: encode error
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// Handle content addressable file storage to a keyspace /cache/<keyspace>/<sha256>
func (s *server) cache(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	key := strings.TrimPrefix(r.URL.Path, "/cache")
	fmt.Println(r.Method, key)

	switch m := r.Method; m {
	case http.MethodGet:
		// download
		r, err := s.bucket.NewReader(ctx, key, nil)
		if err != nil {

			if gcerrors.Code(err) == gcerrors.NotFound {
				w.WriteHeader(http.StatusNotFound)
				return nil
			}
			return err
		}
		defer r.Close()

		if _, err := io.Copy(w, r); err != nil {
			return err
		}
		return nil

	case http.MethodPut:
		// upload
		defer r.Body.Close()

		w, err := s.bucket.NewWriter(ctx, key, nil)
		if err != nil {
			return err
		}
		defer w.Close()

		if _, err := io.Copy(w, r.Body); err != nil {
			return err
		}
		return nil

	case http.MethodHead:
		// info
		ok, err := s.bucket.Exists(ctx, key)
		if err != nil {
			return err
		}
		if !ok {
			w.WriteHeader(http.StatusNotFound)
		}
		return nil

	case http.MethodDelete:
		// delete
		w.WriteHeader(http.StatusNotImplemented)
		return nil

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return nil
	}
}

type file struct {
	svr  *server
	ctx  context.Context
	path string

	r *blob.Reader
	n int64
}

func (f *file) Close() error {
	if f.r != nil {
		return f.r.Close()
	}
	return nil
}

func (f *file) Read(p []byte) (n int, err error) {
	if f.r == nil {
		f.r, err = f.svr.bucket.NewRangeReader(f.ctx, f.path, f.n, -1, nil)
		if err != nil {
			return 0, err
		}
	}
	n, err = f.r.Read(p)
	f.n += int64(n)
	return
}

func (f *file) Seek(offset int64, whence int) (abs int64, err error) {
	switch whence {
	case io.SeekStart:
		abs = offset
	case io.SeekCurrent:
		abs = f.n + offset
	case io.SeekEnd:
		return 0, errors.New("file.Seek: SeekEnd not implemented")
	default:
		return 0, errors.New("file.Seek: invalid whence")
	}
	if abs < 0 {
		return 0, errors.New("file.Seek: negative position")
	}
	if f.n != abs && f.r != nil {
		err = f.Close()
		f.r = nil
	}
	f.n = abs
	return
}

func (f *file) Readdir(count int) (fis []os.FileInfo, err error) {
	dir := path.Dir(f.path)
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	iter := f.svr.bucket.List(&blob.ListOptions{
		Delimiter: "/",
		Prefix:    dir,
	})
	for i := 0; count < 0 || i < count; i++ {
		obj, err := iter.Next(f.ctx)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		mode := os.ModePerm
		if obj.IsDir {
			mode = os.ModeDir
		}
		fis = append(fis, &fileInfo{
			name:    path.Clean(obj.Key),
			size:    obj.Size,
			mode:    mode,
			modtime: obj.ModTime,
		})
	}
	return
}

func (f *file) Stat() (os.FileInfo, error) {
	attr, err := f.svr.bucket.Attributes(f.ctx, f.path)
	if err != nil {
		if gcerrors.Code(err) == gcerrors.NotFound {
			return &fileInfo{
				name: f.path,
				mode: os.ModeDir,
			}, nil
		}
		return nil, err
	}
	return &fileInfo{
		name:    f.path,
		size:    attr.Size,
		mode:    os.ModePerm,
		modtime: attr.ModTime,
	}, nil
}

type fileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modtime time.Time
}

func (fi *fileInfo) Name() string       { return fi.name }
func (fi *fileInfo) Size() int64        { return fi.size }
func (fi *fileInfo) Mode() os.FileMode  { return fi.mode }
func (fi *fileInfo) ModTime() time.Time { return fi.modtime }
func (fi *fileInfo) IsDir() bool        { return fi.mode == os.ModeDir }
func (fi *fileInfo) Sys() interface{}   { return nil }

// Open implements http.FileSystem
func (s *server) Open(name string) (http.File, error) {
	return &file{
		svr:  s,
		ctx:  context.Background(),
		path: name,
	}, nil
}

func (s *server) Serve(l *net.Listener) error {

	m := cmux.New(l)
	httpL := m.Match(cmux.HTTP1Fast())
	grpcL := m.Match(cmux.Any())

	c := make(chan error, 2)
	go func() { c <- s.grpcServer.Serve(grpcL) }()
	defer func() {
		s.grpcServer.GracefulStop()
	}

	go func() { c <- s.httpServer.Serve(httpL) }()
	defer func() {
		s.httpServer.Shutdown(context.Background())
	}

	select {
	case err := <-c:
		return err
	}
}

func NewServer(bucket *blob.Bucket) (*server, error) {
	mux := http.NewServeMux()
	s := &server{
		httpServer: &http.Server{
			Handler: &ochttp.Handler{
				Handler: mux,
			},
		},
		grpcServer: grpc.NewServer(),
		bucket: bucket,
	}

	zpages.Handle(mux, "/debug/")
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/cache/", s.handleErr(s.cache))
	mux.Handle("/", http.FileServer(s))

	exporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		return nil, err
	}
	view.RegisterExporter(exporter)
	mux.Handle("/metrics", exporter)

	return s, nil
}
