package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/bassamelobeid/golang-training/grpc/greeting"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SendGreeting(ctx context.Context, in *pb.Greeting) (*pb.Response, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.Response{Message: "Valar Dohaeris."}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
