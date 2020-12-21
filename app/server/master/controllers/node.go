package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"log"
	"master/models"
	"master/utils"
	"time"
)

type QueryReply struct {
	NodeInfoList []models.NodeInfo `json:"node_info_list"`
}

func QueryNodesAPI(wsConnHandle websocket.Handler) gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.IsWebsocket() {
			wsConnHandle.ServeHTTP(context.Writer, context.Request)
		}
	}
}

func WsConnHandle(conn *websocket.Conn) {


	for {
		queryReply := QueryReply{
			NodeInfoList: models.GetNodeInfoList(),
		}

		if err := websocket.JSON.Send(conn, queryReply); err != nil {
			log.Println(err)
			return
		} else {
			time.Sleep(time.Second)
		}
	}
}

func AddNodeAPI(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "Add Node Success",
	}

	var node models.Node

	if err := c.ShouldBindJSON(&node); err != nil {
		response.Code = 201
		response.Msg = "Parameter Binding Error"
	} else if models.CheckNode(node.IP) {
		response.Code = 202
		response.Msg = "Node Already Exists"
	} else {
		models.AddNode(node)
	}
	utils.StatusOKResponse(response, c)
}

func DeleteNodeAPI(c *gin.Context)  {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "Add Node Success",
	}

	var nodeData models.Node

	if err := c.ShouldBindJSON(&nodeData); err != nil {
		response.Code = 201
		response.Msg = "Parameter Binding Error"
	} else if !models.CheckNode(nodeData.IP) {
		response.Code = 202
		response.Msg = "Node Not Exists"
	} else {
		models.DeleteNode(nodeData.IP)
	}
}
