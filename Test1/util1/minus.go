package main

import (
	"fmt"
	"time"
)

func Minnus(a int,b int) int {
	return a-b
}
func main(){
	exitChan:=make(chan bool,1)
	go func(){
		fmt.Println("等待。。。" )
		time.Sleep(time.Second*2)
		exitChan<-true
		close(exitChan)
	}()

	for{
		b,ok:=<-exitChan
		if !ok{
			fmt.Println("没有值")
			return
		}
		fmt.Println(b)
	}

}