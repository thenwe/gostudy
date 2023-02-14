package main

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/websocket"
	"net/http"

	"context"
	"time"
)

var getConn []*websocket.Conn

/*var connAddr []string*/
var ctx = context.Background()

func main() {
	var upgrader = websocket.Upgrader{
		HandshakeTimeout: 5 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
		//当该端口号和ip地址是第一次访问就存入setWebsocket集合
		v := rdb.SIsMember(ctx, "setWebsocket", conn.RemoteAddr().String()).Val() //判断该地址是否已存在
		if !v {
			err := rdb.SAdd(ctx, "setWebsocket", conn.RemoteAddr().String()).Err() //存储客户端IP地址和端口

			if err != nil {
				fmt.Println(err)
				return
			}
		}
		/*if len(connAddr)==0{
			connAddr=append(connAddr,conn.RemoteAddr().String())
			getConn=append(getConn,conn)
		}
		if len(connAddr)>=1{
			sort.Strings(connAddr)//排序
			if sort.SearchStrings(connAddr,conn.RemoteAddr().String())<0{//判断ip地址是否已存在
				connAddr=append(connAddr,conn.RemoteAddr().String())

			}
		}*/
		getConn = append(getConn, conn)
		go func() {
			for {
				fmt.Println("等待连接")
				msgType, msg, err := conn.ReadMessage()
				if err != nil {
					fmt.Printf("%v连接失败err=%v\n", conn.RemoteAddr(), err)
					//连接失败或者断开就在redis的setWebsocket集合中删除该ip地址和端口号
					_, err = rdb.SRem(ctx, "setWebsocket", 1, conn.RemoteAddr().String()).Result()
					if err != nil {
						fmt.Println(err)
						return
					}
					return
				}
				if len(msg) > 0 {
					fmt.Printf("%s 接收:%s\n", conn.RemoteAddr(), string(msg))
					for i := range getConn {
						err = getConn[i].WriteMessage(msgType, msg)
						if err != nil {
							fmt.Printf("%v发送第%v个失败err=%v\n", conn.RemoteAddr(), i, err)
							continue
						}
					}
				}
			}
		}()
	})
	http.ListenAndServe(":12233", nil)
}
