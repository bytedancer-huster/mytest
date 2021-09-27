package main

import (
	"github.com/gin-gonic/gin"
	"mytest/handler"
)

func main() {
	r := gin.Default()
	r.POST("/create_account", handler.CreateAccount)
	r.POST("/update_password", handler.UpdatePassword)
	r.POST("/user_login", handler.UserLogin)
	r.POST("/user_out", handler.UserOut)
	r.GET("/get_user_info", handler.GetUserInfo)
	r.Run(":80") // listen and serve on 0.0.0.0:8080
}
