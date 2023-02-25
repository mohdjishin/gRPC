package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mohdjishin/gRPC/calculator/calc"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello,I'm a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := calc.NewCalculateServiceClient(conn)

	Calculate(c)

}

func Calculate(c calc.CalculateServiceClient) {
	req := &calc.CalculateRequest{
		Calculating: &calc.Calculating{
			FirstNum:  10,
			SecondNum: 20,
		}}
	res, err := c.Calculate(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Calculate RPC: %v", err)
	}
	log.Printf("Response from Calculate: %v", res.Result)

}
