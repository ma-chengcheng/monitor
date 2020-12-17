package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"log"
	"master/models"
)


type QueryReply struct {
	NodeInfoList []models.NodeInfo `json:"node_info_list"`
}

func Query(wsConnHandle websocket.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.IsWebsocket() {
			wsConnHandle.ServeHTTP(context.Writer, context.Request)
		}
	}
}

type APINode struct {
	IP       string `json:"ip"`
	HostName string `json:"host_name"`
}

func WsConnHandle(conn *websocket.Conn) {

	for {
		queryReply := QueryReply{
			NodeInfoList: models.GetNodeInfoList(),
		}

		if err := websocket.JSON.Send(conn, queryReply); err != nil {
			log.Println(err)
			return
		}
	}
}
