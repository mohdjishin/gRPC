package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mohdjishin/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func main() {

	fmt.Println("Hello World!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v \n", req)
	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()

	result := "Hello, " + firstname + " " + lastname
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v \n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello, " + firstName + " number " + fmt.Sprint(i)

		res := &greetpb.GreetManyTImesResponse{
			Result: result,
		}

		stream.Send(res)
		time.Sleep(time.Second * 2)
	}
	return nil
}
