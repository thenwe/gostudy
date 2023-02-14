package main

import (
	"fmt"
	"runtime"
	"strconv"
	_"time"
)

/*func SayHello(){
	for i:=1;i<=10;i++{
		fmt.Println("hello world")
		time.Sleep(time.Second)
	}
}*/

func main(){
	//获取当前系统cpu的数量
	num :=runtime.NumCPU()
	//设置cpu运行go程序
	runtime.GOMAXPROCS(num)
	fmt.Println("cpu数量"+strconv.Itoa(num))
	/*go SayHello()//协程
	for i:=1;i<10;i++{
		fmt.Println("main:Hello World")
		time.Sleep(time.Second)
	}*/
}

