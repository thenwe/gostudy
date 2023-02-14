package main

import (
	_"bufio"
	"encoding/json"
	"fmt"
	"net"
	_"os"
	_"strings"
)

//客户端
func main(){
	conn,err:=net.Dial("tcp","192.168.73.1:8888")
	defer conn.Close()
	if err!=nil{
		fmt.Println("连接失败，fail:",err)
		return
	}
	for{
		//fmt.Println("连接成功,conn:",conn)
		//功能1：客户端可以发送单行数据 然后就退出
/*		reader:=bufio.NewReader(os.Stdin)//os.Stdin代表标准输入【终端】
		//从终端读取一行用户输入，并准备发送给服务器
		line,err:=reader.ReadString('\n')
		if err!=nil{
			fmt.Println("readString err=",err)
		}
		line=strings.Trim(line,"\r\n")
		if line=="exit"{
			fmt.Println("退出")
			return
		}*/
		//尝试发送切片
		m1:=[]string{"beijing","chengdu"}
		fmt.Println(m1)
		m2,err1:=json.Marshal(m1)
		if err1!=nil{
			fmt.Println("序列化失败")
			return
		}
		n,err:=conn.Write(m2)
		//再将line发送给服务器
		/*n,err:=conn.Write([]byte(line))*/
		if err!=nil{
			fmt.Println("conn.Write err=",err)
		}
		fmt.Printf("客户端发送了%d字节的数据\n",n)
		return
	}

}
