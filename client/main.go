package main

import (
	"log"

	pb "github.com/akhil/grpc-demo-yt/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8081"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fatal %v", err)
	}

	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	names := &namesList{
		Names: []string{"Akhil", "dauren"},
	}

	//callSayHello(client)
	callSayHelloServerStream(client, names)
}
