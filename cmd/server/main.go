package main

import (
	"context"
	"flag"
	"log"
	"net"

	pb "github.com/mohdjishin/gRPC-Secure/internal/api/greeterService"
	"github.com/mohdjishin/gRPC-Secure/pkg/utils"
	"google.golang.org/grpc"
)

// GreeterServer implements the pb.GreeterServer interface.
type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements the SayHello method of the GreeterServer interface.
func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}

func main() {
	// Command line flag to enable mTLS
	mtls := flag.Bool("mtls", false, "Enable mTLS")
	flag.Parse()

	// Load server credentials based on the mtls flag
	var opts []grpc.ServerOption
	if *mtls {
		creds := utils.LoadServerCredentialsWithp12()
		opts = append(opts, grpc.Creds(creds))
	} else {
		opts = append(opts, grpc.Creds(utils.LoadServerCredentials()))
	}

	// Create a new gRPC server with the specified options
	grpcServer := grpc.NewServer(opts...)

	// Register the GreeterServer with the gRPC server
	pb.RegisterGreeterServer(grpcServer, &GreeterServer{})

	// Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening at :50051")

	// Start serving gRPC requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
