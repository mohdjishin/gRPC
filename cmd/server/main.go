package main

import (
	"context"
	"log"
	"net"

	pb "github.com/mohdjishin/gRPC-Secure/internal/api/greeterService"
	"github.com/mohdjishin/gRPC-Secure/pkg/utils"
	"google.golang.org/grpc"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}

func main() {
	grpcServer := grpc.NewServer(grpc.Creds(utils.LoadTLSCredentials()))

	pb.RegisterGreeterServer(grpcServer, &GreeterServer{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
