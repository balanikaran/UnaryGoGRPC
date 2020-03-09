package main

import (
	"fmt"
	"net"
	"context"

	"github.com/krnblni/UnaryGoGRPC/proto"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
)

type greetServer struct{}

func (gs *greetServer) Greet(ctx context.Context, input *greet.Input) (*greet.Output, error) {
	name := input.GetName()
	result := "Hello, " + name + "!"

	return &greet.Output{Greeting: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4500")
	if err != nil {
		fmt.Println("Cannot create a listener...", err)
		return
	}

	grpcGreetServer := grpc.NewServer()

	greet.RegisterGreetServiceServer(grpcGreetServer, &greetServer{})

	if err := grpcGreetServer.Serve(listener); err != nil {
		fmt.Println("Cannot create a server...")
		panic(err)
	}

	fmt.Println("Created Server... ^_^")
}