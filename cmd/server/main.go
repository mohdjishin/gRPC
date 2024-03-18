package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"net"
	"os"

	pb "github.com/mohdjishin/gRPC-Secure/internal/api/greeterService"
	util "github.com/mohdjishin/gRPC-Secure/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GreeterServer implements the greeter service
type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements the SayHello method of the Greeter service
func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}

func main() {
	// Load server certificate and key
	p12File, err := os.ReadFile("certs/server/server.p12")
	if err != nil {
		log.Fatalf("failed to read P12 file: %v", err)
	}
	privKey, leafCert, caCert, err := util.DecodeP12File(p12File, "yourpassword")
	if err != nil {
		log.Fatalf("failed to decode P12 file: %v", err)
	}

	// Convert private key to PEM format
	privKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey.(*rsa.PrivateKey)),
	})

	// Convert leaf certificate to PEM format
	leafCertPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: leafCert.Raw,
	})

	// Create TLS certificate from private key and leaf certificate
	serverCert, err := tls.X509KeyPair(leafCertPEM, privKeyPEM)
	if err != nil {
		log.Fatalf("failed to create TLS certificate: %v", err)
	}

	// Create TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		RootCAs:      caCert,
	})

	// Create a gRPC server with TLS credentials
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	// Register the Greeter service
	pb.RegisterGreeterServer(grpcServer, &GreeterServer{})

	// Start the server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
