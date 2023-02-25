package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mohdjishin/gRPC/calculator/calc"
	"google.golang.org/grpc"
)

type server struct {
	calc.UnimplementedCalculateServiceServer
}

func main() {

	fmt.Println("Calculator Server!z")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calc.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) Calculate(ctx context.Context, req *calc.CalculateRequest) (*calc.CalculateResponse, error) {

	fmt.Printf("Calculate function was invoked with %v", req)

	first_num := req.GetCalculating().GetFirstNum()
	second_num := req.GetCalculating().GetSecondNum()

	result := first_num + second_num
	res := &calc.CalculateResponse{
		Result: result,
	}
	return res, nil

}
