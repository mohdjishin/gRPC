package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/mohdjishin/gRPC/max-val/maxpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello, i'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := maxpb.NewMaxServiceClient(conn)
	fmt.Println(c)
	doMax(c)
}

func doMax(c maxpb.MaxServiceClient) {

	fmt.Println("Starting to do a Client streaming RPC..")

	request := []*maxpb.GetMaxRequest{

		&maxpb.GetMaxRequest{
			Number: 12,
		},
		&maxpb.GetMaxRequest{
			Number: 112,
		},
		&maxpb.GetMaxRequest{
			Number: 1221,
		},
		&maxpb.GetMaxRequest{
			Number: 1512,
		},
		&maxpb.GetMaxRequest{
			Number: 15,
		},
		&maxpb.GetMaxRequest{
			Number: 1,
		},
		&maxpb.GetMaxRequest{
			Number: 10,
		},
	}
	// we create stream by invoking the client
	stream, err := c.GetMax(context.Background())
	if err != nil {

		log.Fatal("Error while creating stream :%v", err)
		return
	}

	// create a channel
	waitc := make(chan struct{})
	// send bunch of message using go routine

	go func() {

		for _, req := range request {
			fmt.Println("sending message :", req)
			stream.Send(req)
			time.Sleep(2 * time.Second)

		}
		stream.CloseSend()

	}()
	// we receive message from the servre
	go func() {

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
	<-waitc
}
