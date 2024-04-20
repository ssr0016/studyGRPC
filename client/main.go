package main

import (
	pb "grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":9000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{
			"John",
			"Jane",
		},
	}

	// callSayHello(client) // it is for unary

	// callSayHelloServerStreaming(client, names) // it is for server streaming

	callSayHelloClientStream(client, names) // it is for client streaming

}
