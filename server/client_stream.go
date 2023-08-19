package main

import (
	pb "go-grpc-learn/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}

		if err != nil {
			log.Fatalf("failed to stream from client %v", err)
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
