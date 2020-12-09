package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "master/message"
	"time"
)

//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080
//}

const address = "127.0.0.1:50051"


func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWatcherClient(conn)

	for true {
		ctx, _ := context.WithTimeout(context.Background(), time.Second * 10)

		r, err := c.GetInfo(ctx, &pb.Request{Code: 200})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Cpu: %v\nMem: %v\nDisk: %v\n", r.GetCpuPercent(), r.GetMemoryPercent(), r.GetDiskPercent())
	}
}
