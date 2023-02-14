package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(time.January, time.Wednesday) //Wednesday

	fmt.Println(time.Now()) //2023-02-14 10:55:22.0067395 +0800 CST m=+0.003690401
	s1 := "成都abc"
	s2 := "成都abC"
	fmt.Println(strings.EqualFold(s1, s2))                            //true判断字符串是否相等 大小写视为相等
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) //返回切片
}
