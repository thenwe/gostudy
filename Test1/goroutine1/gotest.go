package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)
var(
	mymap=make(map[int]int64 ,10)
	lock sync.Mutex
)
func test(n int){
	res:=int64(1)
	for i:=1;i<=n;i++{
		res*=int64(i)
	}

	lock.Lock()
	time.Sleep(time.Second*1)
	mymap[n]=res
	fmt.Println("mymap["+strconv.Itoa(n)+"]:"+strconv.FormatInt(res,10))
	lock.Unlock()
}
func main(){
	for i := 1;i <=200;i++{
		go test(i)

	}

}
