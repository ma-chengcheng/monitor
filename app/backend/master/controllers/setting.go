package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"master/models"
	"master/resource"
	"master/utils"
)

func AddWorker(c *gin.Context) {
	log.Printf("ip: %v\n", c.Param("IP"))

	utils.MysqlDB.Create(&models.Node{
		IP:          c.PostForm("IP"),
		HostName:    c.PostForm("HostName"),
		SSHUsername: c.PostForm("SSHUsername"),
		SSHPassword: c.PostForm("SSHPassword"),
	})

	addSuccess := utils.SubStatus{Code: "add-success", Msg: "添加成功"}
	utils.DefaultResponse(resource.Success, addSuccess.Code, nil, addSuccess.Msg, c)
}
