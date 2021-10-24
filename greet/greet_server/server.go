package main

import (
	"context"
	"fmt"
	"gRPC/greet/greetpb"
	logs "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	//logs "pkg/mod/github.com/sirupsen/logrus@v1.8.1"
	"strconv"
	"time"
)

type Server struct{}

func (*Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with greeting: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName + lastName
	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func (*Server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with greeting: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		res := &greetpb.GreetManyTimesResponse{
			Result: "Hello" + firstName + " number " + strconv.Itoa(i),
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func init() {
	file, _ := os.OpenFile("./logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logs.SetFormatter(&logs.JSONFormatter{})
	logs.SetOutput(file)
}

func main() {
	fmt.Println("Hi From server")
	listener, err := net.Listen("tcp", "127.0.0.1:50051")

	if err != nil {
		logs.WithFields(logs.Fields{
			"Message": "Listening Error",
		}).Warn("Error :", err)
		log.Fatalf("Failed to listen %v", err)
	}

	server := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
