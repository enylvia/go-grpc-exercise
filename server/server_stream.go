package main

import (
	pb "go-grpc-learn/proto"
	"log"
	"time"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameLists, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("got request with names : %v", req.Names)

	for _, v := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + v,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
