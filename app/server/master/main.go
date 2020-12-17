package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"master/resource"
	"master/routers"
)

func main() {
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
	//api.GetNodeInfo("47.101.141.193:50051")
}
