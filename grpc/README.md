# gRPC Example in Go

This folder contains a simple gRPC example implemented in Go, demonstrating client-server communication.

## Project Structure

- `client/`: Contains the gRPC client implementation
- `server/`: Contains the gRPC server implementation
- `greeting/`: Contains the protocol buffer definition

## Prerequisites

- Protocol Buffers compiler
- gRPC Go plugin
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Running the Example

1. Generate the protocol buffer definition:
    ```bash
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/greeting/greeting.proto
    ```

2. Start the server:
   ```bash
   go run server/main.go
   ```
   The server will listen on port 50051 by default.

3. Run the client:
   ```bash
   go run client/main.go
   ```
   By default, the client will connect to `localhost:50051` and send "world" as the name.

### Options

- Server: Use the `-port` flag to specify a different port (e.g., `go run server/main.go -port 8080`)
- Client: 
  - Use the `-addr` flag to specify a different server address (e.g., `go run client/main.go -addr localhost:8080`)
  - Use the `-name` flag to send a different name (e.g., `go run client/main.go -name Alice`)

## Implementation Details

- The server implements a `SayHello` RPC method that responds with "Hello [name]" when called.
- The client sends a name to the server and prints the response.
- Communication is done using gRPC over an insecure connection (for simplicity in this example).

## Protocol Buffer Definition

The `greeting.proto` file defines the service and message types used in this example.

## References

- [gRPC in Go](https://grpc.io/docs/languages/go/quickstart/)
