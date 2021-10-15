package main

import (
	"context"
	"fmt"
	"gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fmt.Println("Hi from client")
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("couldn't connect %v", err)
	}

	defer connection.Close()

	client := greetpb.NewGreetServiceClient(connection)

	doUnary(client)
}

func doUnary(client greetpb.GreetServiceClient) {
	fmt.Println("Starting to do unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "kasra",
			LastName:  "shirazi",
		},
	}
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC: %v", err)
	}
	log.Printf("response from greet %v", res.Result)
}
