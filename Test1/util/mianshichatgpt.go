package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	fmt.Println(x)
}
