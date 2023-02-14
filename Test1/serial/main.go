package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct{
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

func testMap(){
	var a map[string]interface{}
	a=make(map[string]interface{})
	a["name"]="红孩儿"
	a["age"]=30
	a["address"]="洪崖洞"
	//将map序列化
	//将monster序列化
	data,err:=json.Marshal(a)
	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Printf("a序列化之后=%v\n",string(data))
	}
}
func testSlicee(){
	var slice []map[string]interface{}
	m1 := map[string]interface{}{
		"name":"红孩儿",
		"age":12,
		"address":"洪崖洞",
	}
	slice=append(slice, m1)
	data,err:=json.Marshal(slice)
	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Printf("slice序列化之后=%v\n",string(data))
	}
}
func main(){
	monster := Monster{
		Name:"牛魔王",
		Age:500,
		Birthday:"2022-111-11",
		Sal:8000.0,
		Skill:"牛魔拳",
	}

	fmt.Print(monster)
	//将monster序列化
	data,err:=json.Marshal(&monster)
	if err!=nil {
		fmt.Println(err)
	}else{
		fmt.Printf("monster序列化之后=%v\n",string(data))
	}
	testMap()
	testSlicee()
}
