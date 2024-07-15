package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pd "grpc-server/proto"
)

const (
	port = ":50051"
)

type server struct {
	pd.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pd.HelloRequest) (*pd.HelloReply, error) {
	return &pd.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pd.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SayHelloAgain(ctx context.Context, in *pd.HelloRequest) (*pd.HelloReply, error) {
    return &pd.HelloReply{Message: "Hello again " + in.GetName()}, nil
}