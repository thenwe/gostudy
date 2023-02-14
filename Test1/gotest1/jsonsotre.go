package gotest1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct{
	Name string
	Age int
	Skill string
}


func (p *Person) ToJson() bool{
	data , err := json.Marshal(p)
	if err != nil{
		fmt.Printf("序列化失败：%v",err)
	}
	filename := "D:/Go_Study/IoCode/io.txt"
	err=ioutil.WriteFile(filename,data,0666)
	if err != nil{
		fmt.Println("保存出错")
		return false
	}
	return true
}

func ReJson() bool{
	filename := "D:/Go_Study/IoCode/io.txt"
	p:=&Person{}
	data,err:=ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Printf("读文件失败%v",err)
		return false
	}
	err=json.Unmarshal(data,p)
	if err!=nil{
		fmt.Printf("反序列化失败%v",err)
		return false
	}
	fmt.Print(p)
	return true
}
