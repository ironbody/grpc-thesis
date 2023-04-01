package main

import (
	"Benchmarks/grpc/tutorial"
	pb "Benchmarks/grpc/tutorial"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("")
	return &pb.HelloReply{Message: "Hello, " + in.GetName()}, nil
}

func main() {
	println("gRPC server tutorial in Go")

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	tutorial.RegisterGreeterServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
