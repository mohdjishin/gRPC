package utils

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"google.golang.org/grpc/credentials"
	"software.sslmate.com/src/go-pkcs12"
)

func decodeP12File(p12Data []byte, password string) (privateKey interface{}, leaf *x509.Certificate, roots *x509.CertPool, err error) {
	privateKey, certificate, ca, err := pkcs12.DecodeChain(p12Data, password)
	if err != nil {
		return
	}

	leaf = certificate
	roots = x509.NewCertPool()
	for _, intermediateCert := range ca {
		if intermediateCert.IsCA {
			roots.AddCert(intermediateCert)
		}
	}

	return
}

// LoadClientCertificates loads client certificates from a file
func LoadClientCertificates(certFile string) (*tls.Config, error) {
	// Read certificate file
	f, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}

	// Create a certificate pool
	cpool := x509.NewCertPool()
	if !cpool.AppendCertsFromPEM(f) {
		return nil, err
	}

	// Create a TLS configuration with the certificate pool
	cert := &tls.Config{
		RootCAs: cpool,
	}
	return cert, nil
}

// CreateClientCredentials creates client credentials
func CreateClientCredentials(p12FilePath, password string) (credentials.TransportCredentials, error) {
	// Load client certificate and key
	p12Data, err := os.ReadFile(p12FilePath)
	if err != nil {
		return nil, err
	}

	privKey, leafCert, caCert, err := decodeP12File(p12Data, password)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	// Create TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCert,
	})

	return creds, nil
}

// LoadServerCredentials loads server credentials
func LoadServerCredentialsWithp12() credentials.TransportCredentials {
	// Load server certificate and key
	p12File, err := os.ReadFile("certs/server/server.p12")
	if err != nil {
		log.Fatalf("failed to read P12 file: %v", err)
	}
	privKey, leafCert, caCert, err := decodeP12File(p12File, "yourpassword")
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

	return creds
}

func LoadServerCredentials() credentials.TransportCredentials {

	serverCert, err := tls.LoadX509KeyPair("certs/server/server-cert.pem", "certs/server/server-key.pem")
	if err != nil {
		log.Fatalf("Failed to load server certificate and key: %v", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	})

	return creds
}
