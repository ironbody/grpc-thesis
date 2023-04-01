package main

//import (
//	pb "Benchmarks/grpc/tutorial"
//	"context"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//	"log"
//)
//
//const (
//	address     = "localhost:5000"
//	defaultName = "Najem"
//)

// func main() {
// 	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}

// 	defer func(conn *grpc.ClientConn) {
// 		err := conn.Close()
// 		if err != nil {
// 			panic("closed")
// 		}
// 	}(conn)

// 	c := pb.NewGreeterClient(conn)

// 	name := defaultName

// 	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}

// 	log.Printf("Greeting: %s", r.Message)
// }
