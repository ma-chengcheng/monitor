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
	port = "0.0.0.0:50051"
)

type server struct {
	pb.UnimplementedWatcherServer
}

func (s *server) GetInfo(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	log.Printf("[OCTOPUS-node-INFO]\tGETINFO\t%v\n", in.Code)
	return &pb.Reply{
		CpuPercent:    psutil.GetCPUPercent(),
		MemoryPercent: psutil.GetMemPercent(),
		DiskPercent:   psutil.GetDiskPercent(),
	}, nil
}

func main() {
	log.Printf("[OCTOPUS-node-INFO] Listening and serving HTTP on :%v\n", port)
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
