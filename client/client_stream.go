package main

import (
	"context"
	"log"
	"time"

	pb "github.com/akhil/grpc-demo-yt/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Start")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("er")
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("er")
		}
		log.Printf("send")
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("er %v", err)
	}
	log.Printf("%v", res.Messages)
}
