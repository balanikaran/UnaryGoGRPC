package main

import (
	"context"
	"fmt"
	"github.com/krnblni/UnaryGoGRPC/proto"

	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial(":4500", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Unable to create a connection...", err)
		return
	}

	client := greet.NewGreetServiceClient(connection)

	greetRequest := greet.Input{Name: "Karan"}

	greetResponse, err := client.Greet(context.Background(), &greetRequest)
	if err != nil {
		fmt.Println("Error while calling greet...", err)
		return
	}
	fmt.Println("Result: ", greetResponse.GetGreeting())
}