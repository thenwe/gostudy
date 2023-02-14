package testItfImpl

import "fmt"

type ServiceImpl struct {
	Addr string
	Age  int
}

func (service ServiceImpl) GoAway() {
	fmt.Println("goaway")
}
func (service ServiceImpl) ComeAway() {
	fmt.Println("comeaway")
}
