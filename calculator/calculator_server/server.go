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

func (*Server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest,
	stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v", req)
	number := req.GetNumber()

	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to : %v\n", divisor)
		}
	}
	return nil
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
