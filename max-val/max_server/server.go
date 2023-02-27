package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/mohdjishin/gRPC/max-val/maxpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	maxpb.UnimplementedMaxServiceServer
}

func main() {

	fmt.Println("Hello World!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	maxpb.RegisterMaxServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) GetMax(stream maxpb.MaxService_GetMaxServer) error {
	fmt.Println("GetMa function inovoked whith streaming request")

	var max int64
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return nil

		}
		if err != nil {

			log.Fatalf("Error while reading client stream %v", err)
		}
		value := req.GetNumber()
		if max < value {
			max = value
			fmt.Printf("sending new largest value %v \n", max)
			result := "new largest value :" + fmt.Sprintf("%s", max)
			sendErr := stream.Send(&maxpb.GetMaxResponse{
				Result: result,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client %v", err)

			}
		}

	}

}
