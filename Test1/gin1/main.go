package main

import (
	"Test1/gin1/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r1 := r.Group("v1")
	{
		r1.GET("toGetJson", service.ToJson)
		r1.GET("register", service.ToRegister)
		r1.GET("update", service.ToUpdate)
	}
	r.Run(":9091")
}
