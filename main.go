package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/memblob"
	_ "gocloud.dev/blob/s3blob"
)

func env(key, value string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return value
}

var (
	flagAddress = flag.String("address", env("BAZEL_CACHE_ADDRESS", ":8080"), "Address to listen on")
	flagBucket  = flag.String("bucket", env("BAZEL_CACHE_BUCKET", "mem://"), "Bucket storage address")
)

func run() error {
	l, err := net.Listen("tcp", *flagAddress)
	if err != nil {
		return err
	}

	bkt, err := blob.OpenBucket(context.Background(), *flagBucket)
	if err != nil {
		return err
	}
	defer bkt.Close()

	s, err := NewServer(bkt)
	if err != nil {
		return err
	}

	log.Printf("Serving %s", l.Addr())
	return s.Serve(l)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
