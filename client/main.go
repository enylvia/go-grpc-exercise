package main

import (
	"log"

	pb "go-grpc-learn/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect because:%v ", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameLists{
		Names: []string{
			"Aditya",
			"Permadi",
			"Alice",
		},
	}

	// callSayHello(client)

	// callSayHelloServerStream(client, names)

	// callSayHelloClientStream(client, names)
	callSayHelloBirdirectionalStream(client, names)
}
