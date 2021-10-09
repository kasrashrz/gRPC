package main

import (
	"fmt"
	"gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct{}

func main() {
	fmt.Println("Hi From server")
	listener, err := net.Listen("tcp", "127.0.0.1:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
