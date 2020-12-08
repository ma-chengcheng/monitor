package main

import (
	"context"
	"github.com/mamachengcheng/octopus/app/backup/master/master/message"
	"log"
	"net"
)

//import (
//	"context"
//	"log"
//	"net"
//
//	//"google.golang.org/grpc"
//)
//
const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedWatcherServer
}

func (s *server) GetInfo(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
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