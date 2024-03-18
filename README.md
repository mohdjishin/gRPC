# gRPC

This project demonstrates how to implement secure communication between a gRPC server and client using mTLS (Transport Layer Security). Others might add soon.


## Requirements

- Go (Golang)
- Protobuf compiler (`protoc`)
- Go gRPC library
- OpenSSL (for generating certificates)

## Installation



### Running the Server

1. Navigate to the `cmd/server` directory:

    ```bash
    cd cmd/server
    ```

2. Run the server application:

    ```bash
    go run main.go
    ```

### Running the Client

1. Navigate to the `cmd/client` directory:

    ```bash
    cd cmd/client
    ```

2. Run the client application:

    ```bash
    go run main.go
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
