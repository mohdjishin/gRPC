package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/mohdjishin/gRPC-Secure/internal/api/greeterService"
	"github.com/mohdjishin/gRPC-Secure/pkg/utils"
)

func main() {
	creds, err := utils.CreateClientCredentials("certs/client/client.p12", "yourpassword")
	if err != nil {
		log.Fatalf("failed to create client credentials: %v", err)
	}

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Jish"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}
