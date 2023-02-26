package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/mohdjishin/gRPC/squareRoot/sqrroot"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	sqrroot.UnimplementedCalculateServiceServer
}

func main() {

	fmt.Println("Hello World!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	sqrroot.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) SquareRoot(ctx context.Context, req *sqrroot.SquareRootRequest) (*sqrroot.SquareRootResponse, error) {

	fmt.Println("Recived Square root RPC")
	number := req.GetNum()
	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Recived_a_Negative_number : %v", number))

	}

	return &sqrroot.SquareRootResponse{
		NumberRoot: float32(math.Sqrt(float64(number))),
	}, nil

}
