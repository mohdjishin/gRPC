package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mohdjishin/gRPC/average/avg"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello,I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect: %v", err)
	}
	defer conn.Close()
	c := avg.NewAvgServiceClient(conn)

	req := []*avg.AvgRequest{

		&avg.AvgRequest{
			Number: 1,
		},
		&avg.AvgRequest{
			Number: 3,
		}, &avg.AvgRequest{
			Number: 2,
		},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {

		log.Fatalf("Error while calling Avg RPC: %v", err)
	}
	for _, req := range req {
		fmt.Printf("Sending req: %s\n", req.String())
		stream.Send(req)
		time.Sleep(time.Second * 2)

	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v", err)
	}
	fmt.Printf("Avg Response: %v", res.GetResult())
}
