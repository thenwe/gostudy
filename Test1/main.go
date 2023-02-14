package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
type Acccount struct{
	Numcount int
	Encount int
	Spacecount int
	Othercount int
}
func main(){
	filename := "E:/Best/Ber.txt"
	file,err := os.Open(filename)
	if err!=nil{
		fmt.Printf("open file error%v\n",err)
	}
	defer file.Close()
	var account Acccount
	//创建一个reader
	reader := bufio.NewReader(file)
	//开始循环读取filename的内容
	for {
			str,err := reader.ReadString('\n')
			//读到文件末尾
			if err == io.EOF{
				break
			}
			for _,v  := range str{
				switch {
				case v >= 'a' && v <= 'z':
					fallthrough//穿透
				case v >= 'A' && v<='Z':
					account.Encount += 1
				case v == ' ' || v == '\t':
					account.Spacecount++
				case v >= '0' && v <= '9':
					account.Numcount++
				default:
					account.Othercount++
				}
				fmt.Print(string(v))
			}
	}
	fmt.Print("\n")
	fmt.Printf("英文字母的个数:%v,空格的个数:%v,数字的个数:%v",account.Encount,account.Spacecount,account.Numcount)



}
