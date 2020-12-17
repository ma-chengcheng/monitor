package controllers

import (
	"github.com/gin-gonic/gin"
	"master/models"
	"master/utils"
)

func Login(c *gin.Context) {

	var user models.User

	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "Login.vue Success",
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		response.Code = 201
		response.Msg = ""
	} else if !models.CheckUser(user) {
		response.Code = 202
		response.Msg = ""
	} else if token, err := utils.GenerateToken(user.Username, user.Password); err != nil {
		response.Code = 203
		response.Msg = ""
	} else {
		response.Data.(map[string]interface{})["token"] = token
	}

	utils.StatusOKResponse(response, c)
}

func Register(c *gin.Context) {

	var user models.User

	response := utils.Response{
		Code: 200,
		Data: make(map[string]interface{}),
		Msg:  "Register Success",
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		response.Code = 201
		response.Msg = ""
	} else if models.CheckUser(user) {
		response.Code = 202
		response.Msg = ""
	} else {
		models.AddUser(user)
	}

	utils.StatusOKResponse(response, c)
}
