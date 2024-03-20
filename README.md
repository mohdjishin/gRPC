# gRPC

This project demonstrates how to implement secure communication between a gRPC server and client using mTLS (Transport Layer Security). Others might add soon.


## Requirements

- Go (Golang)
- Protobuf compiler (`protoc`)
- Go gRPC library
- OpenSSL (for generating certificates)



## Generating Certificates
Before running the server and client, ensure you have generated the required TLS certificates using OpenSSL.

### Generating Certificates with OpenSSL
- *Run the Certificate Generation Script**: Execute the provided script (`gen.sh`) to generate TLS certificates using OpenSSL. This script will generate certificates and convert them into PKCS#12 format (`.p12` files).
  ```
  chmod +x gen.sh
  ./gen.sh
  ```

### Running the Server

-  Run the server application:

    ```bash
    go run cmd/server/main.go   or     go run cmd/server/main.go -mtls
    ```

### Running the Client

- Run the client application:

    ```bash
    go run cmd/client/main.go   or     go run cmd/client/main.go -mtls
    ```



## Generating Certificates

Before running the server and client, ensure you have generated the required TLS certificates. You can use OpenSSL or any other tool for this purpose. Place the generated certificates(p12) in the appropriate directories (`certs/server` for server certificates(p12) and `certs/client` for client certificates).

## Directory Structure

- `cmd/server`: Contains the server application.
- `cmd/client`: Contains the client application.
- `pkg/`: Contains shared packages and utilities.
- `internal/api/`: Contains Protocol Buffer (.proto) files for defining gRPC services and messages.
- `certs/`: Contains TLS certificates for secure communication.
    - `server/`: Server certificates.
    - `client/`: Client certificates.
