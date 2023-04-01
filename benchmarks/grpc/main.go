package main

import (
	"Benchmarks/grpc/experiment"
	pb "Benchmarks/grpc/experiment"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedExperimentServer
}

var testItem pb.Item = pb.Item{
	Id:            "641eeb68f722093cd36746cd",
	Index:         0,
	Guid:          "8c30a3df-7a8c-4c75-904d-6c8464db486a",
	IsActive:      false,
	Balance:       1569.64,
	Picture:       "http://placehold.it/32x32",
	Age:           22,
	EyeColor:      "blue",
	Name:          "Jeannie Hodge",
	Gender:        "female",
	Company:       "ATGEN",
	Email:         "jeanniehodge@atgen.com",
	Phone:         "+1 (995) 400-2899",
	Address:       "487 Beverley Road, Ahwahnee, Northern Mariana Islands, 1368",
	About:         "Minim do nisi officia nisi nisi velit est. Labore est ex enim qui Lorem aliqua. Aliqua officia reprehenderit minim duis laborum amet occaecat.",
	Registered:    "2020-06-27T08:59:26 -02:00",
	Latitude:      31.12956,
	Longitude:     -10.251515,
	Greeting:      "Hello, Jeannie Hodge! You have 10 unread messages.",
	FavoriteFruit: "strawberry",
}

var test2Return []*pb.Item
var test3Return []*pb.Item

func (s *Server) Test1(ctx context.Context, _ *pb.Empty) (*pb.Test1Res, error) {
	return &pb.Test1Res{Message: "Hello, World!"}, nil
}

func (s *Server) Test2(ctx context.Context, _ *pb.Empty) (*pb.Test23Res, error) {
	return &pb.Test23Res{Items: test2Return}, nil
}

func (s *Server) Test3(ctx context.Context, _ *pb.Empty) (*pb.Test23Res, error) {
	return &pb.Test23Res{Items: test3Return}, nil
}

func main() {
	println("gRPC server tutorial in Go")

	test2Return = make([]*pb.Item, 150)
	for i := 0; i < 150; i++ {
		test2Return[i] = &testItem
	}

	test3Return = make([]*pb.Item, 4500)
	for i := 0; i < 4500; i++ {
		test3Return[i] = &testItem
	}

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	tutorial.RegisterExperimentServer(s, &Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
