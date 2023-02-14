package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:12233/websocket"
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	defer c.Close()
	if err != nil {
		fmt.Println("连接失败:", err)
	}
	fmt.Println("响应成功")
	//done:=make(chan struct{})
	err = c.WriteMessage(websocket.TextMessage, []byte("你好，我是测试人"))
	if err != nil {
		fmt.Println("写入失败:", err)
	}
	for {
		_, message, err := c.ReadMessage()
		fmt.Println(1)
		if err != nil {
			fmt.Println("读返回消息失败:", err)
			break
		}
		fmt.Println("收到消息，", string(message))
	}
}
