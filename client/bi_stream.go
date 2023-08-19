package main

import (
	"context"
	pb "go-grpc-learn/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBirdirectionalStream(client pb.GreetServiceClient, names *pb.NameLists) {

	log.Printf("Birdrectional streaming started")

	stream, err := client.SayHelloBirdStreaming(context.Background())

	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
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
		close(waitc)

	}()

	for _, v := range names.Names {
		req := &pb.HelloRequest{
			Name: v,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error when sending the request %v", err)
		}

		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()

	<-waitc

	log.Printf("Birdirectional streaming finished")

}
