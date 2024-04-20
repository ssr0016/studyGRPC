package main

import (
	"context"

	pb "grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.MessageList, error) {
	return &pb.MessageList{
		Messages: "Hello",
	}, nil
}
