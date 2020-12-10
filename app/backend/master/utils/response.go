package utils

import (
	"github.com/gin-gonic/gin"
)

type SubStatus struct {
	Code string
	Msg  string
}

type response struct {
	Code string      `json:"code"` // 业务返回码
	Data interface{} `json:"data"` //业务返回数据
	Msg  string      `json:"msg"`  // 业务返回描述
}


func DefaultResponse(code int, subCode string, data interface{}, msg string, c *gin.Context) {
	c.JSON(code, response{
		Code: subCode,
		Data: data,
		Msg:  msg,
	})
}
