package main

import (
	"fmt"
	"time"
)

// 无缓冲的管道需要同时接收和发送
func main() {
	st := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()

	ch <- true // 无缓冲，发送方阻塞直到接收方接收到数据。
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
	//time.Sleep(time.Second * 5)
}
