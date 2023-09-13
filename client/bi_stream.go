package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/akhil/grpc-demo-yt/proto"
)

func callHelloBiDirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Direct stre")

	stream, err := client.SayHelloByDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names %v", err)
	}

	log.Printf("%v", stream)

	wait := make(chan chan struct{})

	go func() {
		for {
			message, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("fatal %v", err)
			}

			log.Println(message)
		}

		close(wait)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("fatal")
		}

		time.Sleep(1 * time.Second)
	}

	stream.CloseSend()
	<-wait
	log.Printf("bydirectionally")
}
