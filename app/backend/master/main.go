package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "master/message"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedWatcherServer
}

func (s *server) GetInfo(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Reply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWatcherServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}