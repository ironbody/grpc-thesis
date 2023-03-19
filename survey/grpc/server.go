package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	pb "survey/grpc/protos"
)

type Server struct {
	pb.UnimplementedSenderServer
}

func (s *Server) RetrieveUser(_ context.Context, _ *emptypb.Empty) (*pb.User, error) {
	return &pb.User{
			Name: "Margaret Hamilton",
			Id:   1969,
		},
		nil
}

func main() {
	log.Println("gRPC server started")
	lis, _ := net.Listen("tcp", ":9009")

	s := grpc.NewServer()
	pb.RegisterSenderServer(s, &Server{})
	reflection.Register(s)
	s.Serve(lis)
}
