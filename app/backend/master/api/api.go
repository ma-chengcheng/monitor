package api

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "master/api/message"
	"time"
)

func GetWorkerInfo(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWatcherClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	r, err := c.GetInfo(ctx, &pb.Request{Code: 200})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Cpu: %v\nMem: %v\nDisk: %v\n", r.GetCpuPercent(), r.GetMemoryPercent(), r.GetDiskPercent())
}

func SetWorkerInfo()  {
}

func DeleteWorker()  {

}

func AddWorker()  {

}