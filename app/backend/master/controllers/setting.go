package controllers

import (
	"github.com/gin-gonic/gin"
	"master/models"
	"master/resource"
	"master/utils"
)

func AddWorker(c *gin.Context) {
	node := &models.Node{
		IP:          c.PostForm("IP"),
		HostName:    c.PostForm("HostName"),
		SSHUsername: c.PostForm("SSHUsername"),
		SSHPassword: c.PostForm("SSHPassword"),
	}

	utils.MysqlDB.Create(node)

	addSuccess := utils.SubStatus{Code: "add-success", Msg: "添加成功"}
	utils.DefaultResponse(resource.Success, addSuccess.Code, nil, addSuccess.Msg, c)
}
