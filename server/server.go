package main

import (
	"context"
	"echo-grpc/proto/echopb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const portNumber = "9001"

type server struct{
	echopb.EchoServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("Fail to listen: %v\n", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	echopb.RegisterEchoServiceServer(s, &server{})

	log.Printf("start grpc server on %s port", portNumber)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

func (*server) Echo(ctx context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	message := req.GetMessage()
	fmt.Printf("Echo %v\n", message)

	return &echopb.EchoResponse{
		Result: "Hello " + message,
	}, nil
	
}
