package main

import (
	"fmt"
	sync2 "sync"
)

func main(){
	ch:=make(chan int ,1)
	ch1:=make(chan int ,1)
	var sync sync2.WaitGroup
	sync.Add(2)
	go func() {
		defer sync.Done()
		for i:=1;i<=10;i++{
			select {
			case ch<-1:
				fmt.Println("a")
				ch1<-1
			}
		}
	}()
    go func() {
		defer sync.Done()
    	for i:=1;i<10;i++{
    		select {
			case <-ch1:
				fmt.Println("b")
				<-ch
			}
		}
    }()
	sync.Wait()
}