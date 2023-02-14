package main

import "fmt"

func startChan1(a chan int){
	for i:=1;i<=2000;i++{
		a<-i
	}
	close(a)
}

func main(){
	start1:=make(chan int,1500)
	start2:=make(chan int,4)
	k:=0
	v:=0
	//start3:=make(chan int,4)
	go startChan1(start1)
	var(
		n1,n2,n3,n4 int
		ok1,ok2,ok3,ok4 bool
	)
	go func(){
		for{
			var res int
			res,ok1=<-start1
			n1+=res
			fmt.Println("n1:",n1)
			if !ok1{
				start2<-n1
				/*fmt.Println("协程1结束")
				fmt.Println("写入start1-1" )*/
				break
			}
		}
	}()
	go func(){
		for{
			var res int
			res,ok2=<-start1
			n2+=res
			fmt.Println("n2:",n2)
			if !ok2{
				start2<-n2
				/*fmt.Println("协程2结束")
				fmt.Println("写入start2-1" )*/
				break
			}
		}
	}()
	go func(){
		for{
			var res int
			res,ok3=<-start1
			n3+=res
			fmt.Println("n3:",n3)
			if !ok3{
				start2<-n3
				/*fmt.Println("写入start3-1" )
				fmt.Println("协程3结束")*/

				break
			}
		}
	}()
	go func(){
		for{
			var res int
			res,ok4=<-start1
			n4+=res
			fmt.Println("n4:",n4)
			if !ok4{
				start2<-n4
				/*fmt.Println("协程4结束")
				fmt.Println("写入start4-1" )*/
				break
			}
		}
	}()
	for{
		v+=<-start2
		fmt.Println("--------",v)
		if k==3{
			close(start2)
			fmt.Println("结束")
			break
		}
		k++
	}
}