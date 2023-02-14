package main

import (
	"fmt"

)

func sumCh(sumchan chan int){
	for i:=1;i<=100;i++{
		fmt.Println("写入",i)
		sumchan<-i
	}
	close(sumchan)
}

func isExit(sumchan chan int,boolChan chan bool){
	for {
		str,ok:=<-sumchan
		fmt.Println("读取",str)
		if !ok{
			break
		}
	}
	boolChan<-true
	close(boolChan)
}

func main(){
	sumChan:=make(chan int,100)
	exitChan:=make(chan bool,1)
	go sumCh(sumChan)
	go isExit(sumChan,exitChan)
	for {
		fmt.Println("一次循环")
		b,ok:=<-exitChan
		fmt.Println("循环后的值",b)
		if !ok{
			fmt.Println("退出")
			break
		}
		fmt.Println("读取到数据",b)
	}


}
