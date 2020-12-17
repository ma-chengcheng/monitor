package middlewares

import (
	"github.com/gin-gonic/gin"
	"master/resource"
	util "master/utils"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = resource.Success
		token, _ := c.Cookie("token")

		if token == "" {
			code = resource.ParameterError
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = resource.AuthError
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = resource.AuthError
			}
		}

		if code != resource.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "",
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
