package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "survey/grpc/protos"
)

func main() {
	conn, _ := grpc.Dial("localhost:7070", grpc.WithTransportCredentials(insecure.NewCredentials()))

	defer func(conn *grpc.ClientConn) {
		conn.Close()
	}(conn)

}
