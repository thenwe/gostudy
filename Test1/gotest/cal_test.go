package gotest

import (

	"fmt"
	"testing"
)

func TestHello(t *testing.T){
	res := Add(1,2)

	if res != 3 {
		t.Fatalf("add执行错误，期望值=%v 实际值=%v",8,res)
	}
	t.Logf("add执行正确")
fmt.Print("123")
}
