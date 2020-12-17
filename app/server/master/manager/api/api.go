package api

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "master/manager/message"
	"time"
)

func GetNodeInfo(address string) {

	if conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock()); err == nil {
		c := pb.NewWatcherClient(conn)
		defer conn.Close()

		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

		for {
			if response, err := c.GetInfo(ctx, &pb.Request{Code: 200}); err == nil {
				memoryPercent := response.GetMemoryPercent()
				diskPercent := response.GetDiskPercent()
				cpuPercent := response.GetCpuPercent()
				log.Printf("%v %v %v", memoryPercent, diskPercent, cpuPercent)
			}
			time.Sleep(time.Second)
		}
	}
}