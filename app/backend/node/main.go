package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "node/message"
	"node/psutil"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedWatcherServer
}

func (s *server) GetInfo(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("Received: %v", in.GetCode())
	return &pb.Reply{
		CpuPercent:    psutil.GetCPUPercent(),
		MemoryPercent: psutil.GetMemPercent(),
		DiskPercent:   psutil.GetDiskPercent(),
	}, nil
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
