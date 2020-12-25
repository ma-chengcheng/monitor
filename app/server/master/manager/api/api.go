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
	Status        bool `json:"status"`
}

func GetNodeInfo(address string) string {
	memoryPercent := 0.0
	diskPercent := 0.0
	cpuPercent := 0.0
	status := false

	if conn, err := grpc.Dial(address+":50051", grpc.WithInsecure()); err == nil {
		c := pb.NewWatcherClient(conn)
		defer conn.Close()

		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
		clientDeadline := time.Now().Add(3 * time.Second)
		ctx, cancel := context.WithDeadline(ctx, clientDeadline)
		defer cancel()

		if response, err := c.GetInfo(ctx, &pb.Request{Code: 200}); err == nil {
			memoryPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", response.GetMemoryPercent()), 64)
			diskPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", response.GetDiskPercent()), 64)
			cpuPercent, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", response.GetCpuPercent()), 64)
			status = true
		}
	}

	nodeInfo := NodeInfo{
		MemoryPercent: memoryPercent,
		DiskPercent:   diskPercent,
		CpuPercent:    cpuPercent,
		Status: status,
	}
	result, err := json.Marshal(nodeInfo)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return string(result)
}
