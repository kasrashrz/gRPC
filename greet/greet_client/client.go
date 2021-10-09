package main

import (
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
	fmt.Println(client)
}
