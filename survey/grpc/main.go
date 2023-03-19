package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"reflect"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "survey/grpc/protos"
)

const num = 32

const expectedNum = num * 2

var expectedUser = pb.User{
	Name: "Margaret Hamilton",
	Id:   1969,
}

type CheckerServer struct {
	pb.UnimplementedCheckerServer
}

func (s *CheckerServer) GetNumber(_ context.Context, _ *emptypb.Empty) (*pb.NumberReply, error) {
	log.Printf("Sent number: %d\n", num)
	return &pb.NumberReply{Number: num}, nil
}

func (s *CheckerServer) CheckNumber(_ context.Context, in *pb.CheckNumberRequest) (*pb.CheckNumberReply, error) {
	log.Printf("Checking number: %d\n", in.Number)
	if in.Number == expectedNum {
		log.Printf("%d was correct!", in.Number)
		return &pb.CheckNumberReply{Message: "Correct!"}, nil
	}

	log.Printf("%d was incorrect!", in.Number)
	return &pb.CheckNumberReply{Message: "Incorrect!"}, nil

}

func (s *CheckerServer) CheckUser(_ context.Context, _ *emptypb.Empty) (*pb.CheckUserReply, error) {
	conn, _ := grpc.Dial("localhost:9009", grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer func(conn *grpc.ClientConn) {
		conn.Close()
	}(conn)

	c := pb.NewSenderClient(conn)

	result, _ := c.RetrieveUser(context.Background(), &emptypb.Empty{})

	log.Println(result)
	log.Println(&expectedUser)
	log.Println(reflect.DeepEqual(result, &expectedUser))

	if reflect.DeepEqual(result, &expectedUser) {
		return &pb.CheckUserReply{Message: "Correct!"}, nil
	}

	return &pb.CheckUserReply{Message: "Incorrect"}, nil
}

func main() {
	log.Println("gRPC server started")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCheckerServer(s, &CheckerServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
