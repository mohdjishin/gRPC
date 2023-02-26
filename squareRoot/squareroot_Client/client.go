package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mohdjishin/gRPC/squareRoot/sqrroot"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {

	fmt.Println("Hello, i'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := sqrroot.NewCalculateServiceClient(conn)
	fmt.Println(c)

	doErrorUnary(c)
}

func doErrorUnary(c sqrroot.CalculateServiceClient) {
	fmt.Println("Starting to do  square root Unary rcp")

	res, err := c.SquareRoot(context.Background(), &sqrroot.SquareRootRequest{Num: 10})

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// ok means actual error from grpc
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("we  probably sent neg number!")
			}
		} else {
			log.Fatalf("Big error calling squareRoot %v", err)
		}
	}

	fmt.Printf("Result of square root of %v: %v ", 4, res.GetNumberRoot())
}
