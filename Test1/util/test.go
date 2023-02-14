package main

import (
	"fmt"
	"time"
)

func main(){
	ch1:=make(chan int,2)
	ch2:=make(chan int,4)
	ch1<-1
	ch1<-1
	go func(){
		for{
			fmt.Println("1")
			select {
			case <-ch1:
				time.Sleep(time.Microsecond*10)
				fmt.Println("线程1获取ch1")
			default:
				ch2<-1
				fmt.Println("线程1获取ch2")
				time.Sleep(time.Microsecond)
				break
			}
		}
	}()
	go func(){
		for{
			fmt.Println("1")
			select {
			case <-ch1:
				time.Sleep(time.Microsecond*10)
				fmt.Println("线程2获取ch1")
			default:
				ch2<-1
				fmt.Println("线程2获取ch2")
				time.Sleep(time.Microsecond)
				break
			}
		}
	}()
	go func(){
		for{
			fmt.Println("1")
			select {
			case <-ch1:
				time.Sleep(time.Microsecond*10)
				fmt.Println("线程3获取ch1")
			default:
				ch2<-1
				fmt.Println("线程3获取ch2")
				time.Sleep(time.Microsecond)
				break
			}
		}
	}()
	go func(){
		for{
			fmt.Println("1")
			select {
			case <-ch1:
				time.Sleep(time.Microsecond*10)
				fmt.Println("线程4获取ch1")
			default:
				ch2<-1
				fmt.Println("线程4获取ch2")
				time.Sleep(time.Microsecond)
				break
			}
		}
	}()

	for{
		//fmt.Println("程序进行中")
		if len(ch2)==4{
			fmt.Println("程序结束")
			close(ch2)
			close(ch1)
			break
		}
	}
}
