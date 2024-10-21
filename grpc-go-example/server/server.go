package main

import (
	"context"
	"log"
	"net"

	pb "grpc-go-example/protos"

	"google.golang.org/grpc"
)

// server is used to implement ExampleService.
type server struct {
	pb.UnimplementedExampleServiceServer
}

// SayHello implements ExampleService
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExampleServiceServer(s, &server{})
	log.Printf("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
