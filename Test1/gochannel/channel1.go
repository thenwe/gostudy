package main

import "fmt"

func main(){
	intChan :=make(chan int,3)
	intChan<-100
	intChan<-200
	close(intChan)//关闭管道，这时不能再向管道写入数据
	fmt.Println("写入成功")
	//当管道关闭后，读取数据是可以的
	n1:=<-intChan
	fmt.Println("n1=",n1)
	intChan1:=make(chan int ,200)
	for i:=1;i<=100;i++{
		intChan1<-i*2
	}
	//在遍历管道之前必须先关闭管道
	close(intChan1)
	for v:= range intChan1{
		fmt.Println("v=",v)
	}

}
