package main

import (
	"context"
	"groupcachepb"
)

// key == "file_offset"

func (s *server) Get(ctx context.Context, req *groupcachepb.GetRequest) (*groupcachepb.GetResponse, error) {
	return &groupcachepb.GetResponse{
		Value:     []byte{},
		MinuteQps: 0,
	}, nil
}
