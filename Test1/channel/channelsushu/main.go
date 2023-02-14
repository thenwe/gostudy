package main

import "fmt"

var t =0
func all(a chan int){
	for i:=1;i<=100;i++{
		a<-i
		//fmt.Println("存储的数据",i)
	}
	close(a)
}
func isSu(a chan int,b chan int){
	for{
		v,ok:=<-a

		for i:=1;i<=v/2;i++{
			if v%i==0 && i!=1 {
				break
			}
			if i==v/2{
				fmt.Println("协程1")
				b<-v
			}
		}
		if !ok{
			fmt.Println("------v:",v)
			t++
			break
		}

	}
}
func isSu1(a chan int,b chan int){
	for{
		v,ok:=<-a
		if !ok{
			t++
			break
		}
		for i:=1;i<=v/2;i++{
			if v%i==0 && i!=1 {
				break
			}
			if i==v/2{
				fmt.Println("协程2")
				b<-v
			}
		}

	}
}
func main(){
	suChan:=make(chan int,100)
	isChan:=make(chan int,100)
	go all(suChan)
	go isSu(suChan,isChan)
	go isSu1(suChan,isChan)
	for{
		v,ok:=<-isChan
		fmt.Println("素数",v,ok)
		if t==2 && len(isChan)==0{
			close(isChan)
			fmt.Println("结束")
			break
		}

	}
}