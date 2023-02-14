package main

import (
	"fmt"
	"io"
	"net"
)
func process(c net.Conn){
	//循环接收客户端发送的数据
	defer c.Close()
	//var getSlice []string
	for{
		//创建一个新的切片
		buf:=make([]byte,1024)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write[发送]，那么协程就会阻塞在这里
		fmt.Println("服务器在等待客户端发送信息"+c.RemoteAddr().String())
		n,err:=c.Read(buf)//从conn读取
		if err==io.EOF{
			fmt.Println("客户端退出")
			return
		}
		//3.显示客户端发送的内容到服务器的终端
		/*err2:=json.Unmarshal(buf[:n],&getSlice)//将byte切片反序列化为字符串切片
		if err2!=nil{
			fmt.Println("反序列化失败")
		}*/
		fmt.Println(string(buf[:n]))
		//fmt.Println(getSlice)
	}
}


func main(){

	fmt.Println("服务器等待连接")
	listen,err:=net.Listen("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println("listen err=",err)
		return
	}
	defer listen.Close()//延时关闭listen
	//循环等待客户端来连接
	for{
		fmt.Println("等待客户端来连接")
		conn,err:=listen.Accept()//阻塞
		if err!=nil{
			fmt.Println("Accept() err=",err)
			continue
		}else{
			fmt.Printf("Accept() suc con=%v\n,客户端ip=%v\n",conn,conn.RemoteAddr().String())
			go process(conn)
		}
		//开启协程，为客户端服务
	}

	//fmt.Printf("listen suc=%v\n",listen)
}