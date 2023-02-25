package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/mohdjishin/gRPC/average/avg"
	"google.golang.org/grpc"
)

type server struct {
	avg.UnimplementedAvgServiceServer
}

func main() {

	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		fmt.Println("Error in listening")
	}

	s := grpc.NewServer()

	avg.RegisterAvgServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("Error in serving")
	}

}

func (*server) Avg(stream avg.AvgService_AvgServer) error {
	fmt.Println("Avg function was invoked with a streaming request")
	var sum int64
	count := 0
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&avg.AvgResponse{
				Result: sum / int64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		num := msg.GetNumber()

		count = count + 1
		sum = sum + num
		fmt.Println("Received number: ", num)
	}
}
