package main

import (
	"log"
	"net"

	pb "github.com/akhil/grpc-demo-yt/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("service", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to start %v", err)
	}
}
