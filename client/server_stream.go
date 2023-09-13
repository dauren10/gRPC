package main

import (
	"context"
	"io"
	"log"

	pb "github.com/akhil/grpc-demo-yt/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Hello")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not same names %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while %v", err)
		}

		log.Println(message)
	}

	log.Printf("Streaming finished %v", stream)
}
