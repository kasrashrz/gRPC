package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("Hi From Calculator Client")
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("couldn't connect %v", err)
	}

	defer connection.Close()

	client := calculatorpb.NewCalculatorServiceClient(connection)

	doUnary(client)
}

func doUnary(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	}

	res, err := client.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling calculator RPC: %v", err)
	}
	log.Printf("response from calculator %v", res.SumResult)
}
