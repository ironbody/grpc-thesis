package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	pb "survey/grpc/protos"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer func(conn *grpc.ClientConn) {
		conn.Close()
	}(conn)

	c := pb.NewCheckerClient(conn)

	num, _ := c.GetNumber(context.Background(), &emptypb.Empty{})

	log.Printf("Number: %v", num.Number)

	num.Number *= 2

	r, _ := c.CheckNumber(context.Background(), &pb.CheckNumberRequest{Number: num.Number})

	log.Println(r)

	/*
		Assuming the function returns an object
		with the following fields:
		Name string
		Friends int32
	*/

	r.Message = "hey"

}
