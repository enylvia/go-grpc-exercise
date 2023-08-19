package main

import (
	"context"
	pb "go-grpc-learn/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameLists) {
	log.Printf("Streaming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("some error occured when sending names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}

		log.Println(message)
	}

	log.Printf("Streaming finished")
}
