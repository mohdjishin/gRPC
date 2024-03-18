package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/mohdjishin/gRPC-Secure/internal/api/greeterService"
	"github.com/mohdjishin/gRPC-Secure/pkg/utils"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	useMTLS := flag.Bool("mtls", false, "Use mutual TLS authentication")
	flag.Parse()

	var dialOpt grpc.DialOption
	var err error

	fmt.Println("mtls flag value:", *useMTLS)
	if *useMTLS {
		tlsConfig, err := utils.CreateClientCredentials("certs/client/client.p12", "yourpassword")
		if err != nil {
			return err
		}
		dialOpt = grpc.WithTransportCredentials(tlsConfig)
	} else {
		tlsConfig, err := utils.LoadClientCertificates("certs/CA/ca-cert.pem")
		if err != nil {
			log.Println("Error loading client certificates:", err)
			return err
		}
		dialOpt = grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))
	}

	conn, err := grpc.Dial("0.0.0.0:50051", dialOpt)
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Jish"})
	if err != nil {
		return err
	}

	log.Printf("Greeting: %s", resp.Message)
	return nil
}
