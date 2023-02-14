package main

import (
	itf2 "Test1/itf"
	"Test1/itf/testItfImpl"
	"fmt"
)

func main() {
	var itf itf2.Meth
	impl := testItfImpl.ServiceImpl{Addr: "绵阳", Age: 23}
	fmt.Println(itf == nil)
	itf = impl
	itf.GoAway()
	itf.ComeAway()

}
