package main

import "fmt"

func main(){
	ch1:=make(chan int ,10)
	go func(){
		ch1<-1
		close(ch1)
	}()

	for{
		v:=<-ch1
		fmt.Println("输出v:",v)
	}
}