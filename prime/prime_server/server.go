package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mohdjishin/gRPC/prime/prime"
	"google.golang.org/grpc"
)

type server struct {
	prime.UnimplementedPrimeServiceServer
}

func main() {
	fmt.Println("Hello,I'm a server")
	conn, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	s := grpc.NewServer()
	prime.RegisterPrimeServiceServer(s, &server{})
	if err := s.Serve(conn); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}

func (s *server) Prime(req *prime.PrimeRequest, stream prime.PrimeService_PrimeServer) error {
	fmt.Println("Prime function was invoked with %v \n", req)

	lim := req.GetNumber()

	for i := 0; i < int(lim); i++ {

		if isPrime(i) {

			res := &prime.PrimeResponse{
				Prime: int64(i),
			}
			stream.Send(res)
			time.Sleep(time.Second * 2)

		}
	}

	return nil
}

func isPrime(n int) bool {
	if 1 >= n {
		return false
	}

	for i := 2; i <= n/2; i++ {

		if n%i == 0 {
			return false
		}

	}
	return true
}
