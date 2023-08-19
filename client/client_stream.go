package main

import (
	"context"
	pb "go-grpc-learn/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameLists) {
	log.Printf("client streaming started")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names %v ", err)
	}

	for _, v := range names.Names {
		req := &pb.HelloRequest{
			Name: v,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("could not send stream %v", err)
		}

		log.Printf("Sent the request with name: %s", v)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finished")

	if err != nil {
		log.Fatalf("failed to recieve respond %v", err)
	}

	log.Printf("%v", res.Message)
}
