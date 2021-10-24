package main

import (
	"context"
	"fmt"
	"gRPC/calculator/calculatorpb"
	"google.golang.org/grpc"
	"io"
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

	//doUnary(client)
	doServerStreaming(client)

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

func doServerStreaming(client calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a PrimeDecomposition Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 1234567890,
	}
	stream, err := client.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeDecomposition RPC: %v ", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("something went wrong: %v ", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}
