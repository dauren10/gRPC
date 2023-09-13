package main

import (
	"io"
	"log"

	pb "github.com/akhil/grpc-demo-yt/proto"
)

func (s *helloServer) SayHelloByDirectionalStreaming(stream pb.GreetService_SayHelloByDirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		log.Printf(req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		log.Printf(res.Message)
	}
}
