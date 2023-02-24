package main

import (
	"fmt"
	"log"

	"github.com/mohdjishin/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello,I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client: %f", c)
}
