package main

import (
	"context"
	"fmt"
	"io"
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
	doServerStreaming(c)

}

func doServerStreaming(c prime.PrimeServiceClient) {

	res := &prime.PrimeRequest{Number: 10}

	req, err := c.Prime(context.Background(), res)
	if err != nil {
		log.Fatalf("Error while calling Prime RPC: %v", err)
	}

	for {

		msg, err := req.Recv()

		if err == io.EOF {

			return // we've reached the end of the stream
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
			return
		}

		log.Printf("Response from Prime: %v", msg.GetPrime())
	}

}
