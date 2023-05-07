package main

import (
	pb "Benchmarks/grpc/experiment"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type Server struct {
	pb.UnimplementedExperimentServer
}

var test1Return = pb.Test1Res{Message: "Hello, World!"}
var test2Return pb.Test1Res
var test3Return pb.Test1Res

func (s *Server) Test1(ctx context.Context, _ *pb.Empty) (*pb.Test1Res, error) {
	return &test1Return, nil
}

func (s *Server) Test2(ctx context.Context, _ *pb.Empty) (*pb.Test1Res, error) {
	return &test2Return, nil
}

func (s *Server) Test3(ctx context.Context, _ *pb.Empty) (*pb.Test1Res, error) {
	return &test3Return, nil
}

func main() {
	log.Println("gRPC server tutorial in Go")

	dat1, _ := os.ReadFile("150k.txt")
	dat2, _ := os.ReadFile("3M.txt")

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	log.Printf("150k len: %v", len(string(dat1)))
	log.Printf("3m len: %v", len(string(dat2)))

	test2Return = pb.Test1Res{Message: string(dat1)}

	test3Return = pb.Test1Res{Message: string(dat2)}

	s := grpc.NewServer()
	pb.RegisterExperimentServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
