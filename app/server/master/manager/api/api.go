package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

type NodeInfo struct {
	MemoryPercent float64 `json:"memory_percent"`
	DiskPercent   float64 `json:"disk_percent"`
	CpuPercent    float64 `json:"cpu_percent"`
}

func GetNodeInfo(address string) string {

	//if conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock()); err == nil {
	//	c := pb.NewWatcherClient(conn)
	//	defer conn.Close()
	//
	//	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	//
	//	if response, err := c.GetInfo(ctx, &pb.Request{Code: 200}); err == nil {
	//		memoryPercent := response.GetMemoryPercent()
	//		diskPercent := response.GetDiskPercent()
	//		cpuPercent := response.GetCpuPercent()
	//		log.Printf("%v %v %v %v", address, memoryPercent, diskPercent, cpuPercent)
	//	}
	//}
	log.Printf("ip: %v", address)
	MemoryPercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64() * 100), 64)
	DiskPercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64() * 100), 64)
	CpuPercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", rand.Float64() * 100), 64)

	nodeInfo  := NodeInfo{
		MemoryPercent: MemoryPercent,
		DiskPercent:   DiskPercent,
		CpuPercent:    CpuPercent,
	}
	result, err := json.Marshal(nodeInfo)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return string(result)
}
