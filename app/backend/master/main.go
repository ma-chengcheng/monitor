package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"master/resource"
	"master/routers"
)

func main() {
	cfg, err := ini.Load(resource.ConfFilePath)

	if err != nil {
		panic(err)
	}

	server := cfg.Section("server")

	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	routers.InitRouter(router)

	address := server.Key("http").String() + ":" + server.Key("port").String()
	router.Run(address)
}
