package api

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	pb "master/manager/message"
	"strconv"
	"time"
)

type NodeInfo struct {
	MemoryPercent float64 `json:"memory_percent"`
	DiskPercent   float64 `json:"disk_percent"`
	CpuPercent    float64 `json:"cpu_percent"`
}

func GetNodeInfo(address string) string {
	memoryPercent := 0.0
	diskPercent := 0.0
	cpuPercent := 0.0


	if conn, err := grpc.Dial(address+":50051", grpc.WithInsecure(), grpc.WithBlock()); err == nil {
		c := pb.NewWatcherClient(conn)
		defer conn.Close()

		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

		if response, err := c.GetInfo(ctx, &pb.Request{Code: 200}); err == nil {
			memoryPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", response.GetMemoryPercent()), 64)
			diskPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", response.GetDiskPercent()), 64)
			cpuPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", response.GetCpuPercent()), 64)
		}
	}

	nodeInfo  := NodeInfo{
		MemoryPercent: memoryPercent,
		DiskPercent:   diskPercent,
		CpuPercent:    cpuPercent,
	}
	result, err := json.Marshal(nodeInfo)
	if err != nil {
		log.Fatalf("%v", err)
	} else {
		log.Printf("%v", result)
	}
	return string(result)
}
