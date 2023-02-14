package main

import "fmt"

type Author struct {
	Name         int      `json:Name`
	Publications []string `json:Publication,omitempty`
}
type Set map[string]struct{} //模拟集合

func main() {
	type si[t int | string] []t //泛类型
	type si1[r int | string | si[int]] struct {
		ages si[int]
		addr r
	}
	s1 := si1[si[int]]{
		ages: si[int]{1, 23},
		addr: si[int]{1, 23},
	}
	s2 := si1[int]{
		ages: si[int]{1, 23},
		addr: 123,
	}
	var s si[int] = []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s)
	sample := "我爱go"
	runeSamples := []rune(sample) //强制转换 int 32
	runeSamples[0] = '你'
	fmt.Println(runeSamples)                            //[20320 29233 103 111]
	fmt.Println(string(runeSamples))                    //你爱go
	fmt.Println(string([]rune{20320, 29233, 103, 111})) //你爱go
	fmt.Println(string([]rune{'3', '这'}))               //3这
	fmt.Println([]rune{'3', '这'})                       //[51 36825]
	//判断map中是否包含某个元素
	var sample1 map[int]int
	if _, err := sample1[10]; err {
		fmt.Println("包含10")
	} else {
		fmt.Println("不包含10")
	}
	//结构体打印时，%v 和 %+v 的区别
	var author Author
	author.Name = 2
	author.Publications = []string{"23", "12", "sjsjssj"}
	fmt.Printf("%v:\n", author)
	fmt.Printf("%+v:\n", author)
	fmt.Printf("%#v:\n", author)
	//用空结构体模拟集合
	Set := make(Set)
	for _, v := range []string{"A", "B"} {
		Set[v] = struct{}{}
	}
	if _, err := Set["A"]; err {
		fmt.Println("Set集合包含A")
	} else {
		fmt.Println("Set集合不包含A")
	}
}
