package main

import (
	"gin/db"
	"gin/dto"
	"github.com/gin-gonic/gin"
)

var sUser dto.SocketUser

func main() {
	r := gin.Default()
	r.Use(Cors())
	//r.LoadHTMLGlob("../template/*")
	r.Static("/static", "./static")
	r.GET("register", db.Register)
	//登录
	r.GET("login", db.BeforeLogin, db.Login)
	r.Run(":8081")
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		//c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}
