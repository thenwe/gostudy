package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net"
)

func ViewServer(c *gin.Context) {
	conn,_:=net.Dial("tcp","192.168.73.1:8888")

	data:= c.Query("data")
	v,_:=json.Marshal(data)

	conn.Write(v)

	c.JSON(200,gin.H{
		"message":"成功",
	})

}
