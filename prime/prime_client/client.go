package main

import (
	"fmt"
	"log"

	"github.com/mohdjishin/gRPC/prime/prime"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello i'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := prime.NewPrimeServiceClient(conn)
	fmt.Println(c)
}
