package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/mohdjishin/gRPC/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet function was invoked with a streaming request")
	result := "Hello "
	for {

		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result})

		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err

		}

		firstName := msg.GetGreeting().GetFirstName()
		result += firstName + "!"
		fmt.Println("Hello " + firstName)

	}

}

func (s *server) GreateEveryone(stream greetpb.GreetService_GreateEveryoneServer) error {
	fmt.Println("GreetEveryone function invocked with a streaming request")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {

			log.Fatalf("Error while reading client stream %v", err)
		}
		firstname := req.GetGreeting().GetFirstName()

		result := "Hello " + firstname + "!"
		sendErr := stream.Send(&greetpb.GreetEveryoneResponse{

			Result: result,
		})

		if sendErr != nil {
			log.Fatalf("Error while sending data to client %v", err)

		}

	}

}

func (s *server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("Greet function was invoked with %v \n", req)
	for i := 0; i < 3; i++ {

		if ctx.Err() == context.Canceled {
			fmt.Println("The client canceled the request")
			return nil, status.Error(codes.DeadlineExceeded, "The client canceled the request")
		}
		time.Sleep(1 * time.Second)

	}

	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()

	result := "Hello, " + firstname + " " + lastname
	res := &greetpb.GreetWithDeadlineResponse{
		Result: result,
	}

	return res, nil
}
