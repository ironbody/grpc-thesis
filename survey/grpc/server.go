package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	pb "survey/grpc/protos"
)

type Server struct {
	pb.UnimplementedSenderServer
}

func (s *Server) RetrieveUser(_ context.Context, _ *emptypb.Empty) (*pb.User, error) {

}

func main() {
	log.Println("gRPC server started")
	lis, _ := net.Listen("tcp", ":9009")

	s := grpc.NewServer()
	pb.RegisterSenderServer(s, &Server{})
	s.Serve(lis)
}
