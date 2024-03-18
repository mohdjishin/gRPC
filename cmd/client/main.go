package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/mohdjishin/gRPC-Secure/internal/api/greeterService"
	util "github.com/mohdjishin/gRPC-Secure/pkg/utils"
)

func main() {
	// Load client certificate and key
	p12Data, err := os.ReadFile("certs/client/client.p12")
	if err != nil {
		log.Fatalf("failed to read P12 file: %v", err)
	}

	privKey, leafCert, caCert, err := util.DecodeP12File(p12Data, "yourpassword")
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

	// Create TLS certificate from leaf certificate and private key
	clientCert, err := tls.X509KeyPair(leafCertPEM, privKeyPEM)
	if err != nil {
		log.Fatalf("failed to create TLS certificate: %v", err)
	}

	// Create TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCert,
	})

	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Jish"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}
