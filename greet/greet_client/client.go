package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// fmt.Printf("Created client: %f", c)
	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBiDirStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Jishin",
			LastName:  "Mohd",
		}}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{

		Greeting: &greetpb.Greeting{

			FirstName: "Jishin",
			LastName:  "Mohd",
		}}

	res, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			return // we've reached the end of the stream

		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
			return
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Client Streaming RPC...")

	req := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jishin",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jamal",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Nibras",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Waleed",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}
	//  we iterate over our slice and send each message individually
	for _, r := range req {
		stream.Send(r)
		time.Sleep(time.Second * 1)

	}

	res, err := stream.CloseAndRecv()
	fmt.Println("longGreet request sent", res)

}

func doBiDirStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Client streaming RPC..")

	request := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jishin",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jamal",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Nibras",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Waleed",
			},
		},
	}

	// we create stream by invoking the client
	stream, err := c.GreateEveryone(context.Background())
	if err != nil {
		log.Fatal("Error while creating stream :%v", err)
		return
	}
	waitc := make(chan struct{})
	//  we send a buch of message to the client (go routine)
	go func() {
		// func to send msg
		for _, req := range request {
			fmt.Println("sending message ", req)

			stream.Send(req)
			time.Sleep(time.Second * 1)

		}
		stream.CloseSend()

	}()
	//  we receive a buch of message fromm the server (go routine )
	go func() {
		// recive to send msg
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Error while reciving %v", err)
			}
			fmt.Printf("Recived %v\n", res.GetResult())
		}
	}()
	// block until everything is done
	<-waitc
}
