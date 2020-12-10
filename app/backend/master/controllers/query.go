package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"log"
)

type WorkerInfo struct {
	HostName string  `json:"host_name"`
	IP       string  `json:"ip"`
	Cpu      float64 `json:"cpu"`
	Mem      float64 `json:"mem"`
	Disk     float64 `json:"disk"`
}

type QueryReply struct {
	WorkerInfos []WorkerInfo
}

func Query(wsConnHandle websocket.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.IsWebsocket() {
			wsConnHandle.ServeHTTP(context.Writer, context.Request)
		}
	}
}


func WsConnHandle(conn *websocket.Conn) {
	var name = 1

	for {

		queryReply := QueryReply{
			WorkerInfos: []WorkerInfo{
				{
					HostName: "da",
					IP: "dad",
					Cpu: float64(name),
					Mem: 1,
					Disk: 1,
				},
			},
		}
		if err := websocket.JSON.Send(conn, queryReply); err != nil {
			log.Println(err)
			return
		}
	}
}
