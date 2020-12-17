package controllers

import (
	"github.com/gin-gonic/gin"
	"master/models"
	"master/utils"
)

func AddNode(c *gin.Context) {
	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "Register Success",
	}

	var node models.Node

	if err := c.ShouldBindJSON(&node); err != nil {
		response.Code = 201
		response.Msg = ""
	} else if models.CheckNode(node.IP) {
		response.Code = 202
		response.Msg = ""
	} else {
		models.AddNode(node)
	}
	utils.StatusOKResponse(response, c)
}

func DeleteNode(c *gin.Context)  {
	var nodeData models.Node

	if err := c.ShouldBindJSON(&nodeData); err == nil {
		if models.CheckNode(nodeData.IP) {
			models.DeleteNode(nodeData.IP)
		}
	}
}

