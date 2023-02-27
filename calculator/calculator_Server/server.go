package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mohdjishin/gRPC/calculator/calc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	reflection.Register(s)

	calc.RegisterCalculateServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) Calculate(ctx context.Context, req *calc.CalculateRequest) (*calc.CalculateResponse, error) {

	fmt.Printf("Calculate function was invoked with %v", req)

	first_num := req.GetCalculating().GetFirstNum()
	second_num := req.GetCalculating().GetSecondNum()
	operation := req.GetCalculating().GetOperation()
	var res interface{}
	switch operation {
	case "add":
		res = first_num + second_num
	case "sub":
		res = first_num - second_num
	case "mul":
		res = first_num * second_num
	case "div":
		res = first_num / second_num
	default:
		res = "Invalid Operation"
	}

	result := &calc.CalculateResponse{
		Result: fmt.Sprintf("%v", res),
	}
	return result, nil

}
