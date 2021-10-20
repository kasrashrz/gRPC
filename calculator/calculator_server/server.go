package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct{}

func (*Server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v", req)
	firstNum := req.FirstNumber
	secondNum := req.SecondNumber
	sum := firstNum + secondNum
	response := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return response, nil
}

func main() {
	fmt.Println("Hi From Calculator Server ")
	listener, err := net.Listen("tcp", "127.0.0.1:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
