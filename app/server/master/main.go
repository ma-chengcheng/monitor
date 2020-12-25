package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
	"master/manager/api"
	"master/models"
	"master/resource"
	"master/routers"
)

func ServerAPI()  {
	log.Printf("Start Server API")
	for {
		for _, ip := range models.GetIPList() {
			nodeInfo := api.GetNodeInfo(ip)
			models.SetNodeInfo(ip, nodeInfo)
		}
	}
}

func main() {

	go ServerAPI()

	gin.SetMode(gin.DebugMode)

	if cfg, err := ini.Load(resource.ConfFilePath); err != nil {
		panic(err)
	} else {
		server := cfg.Section("server")
		address := server.Key("http").String() + ":" + server.Key("port").String()

		router := gin.Default()
		routers.InitRouter(router)
		router.Run(address)
	}

}
