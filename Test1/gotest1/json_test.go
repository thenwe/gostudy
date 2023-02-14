package gotest1

import "testing"

func TestJson(t *testing.T){
	p := Person{
		Name:  "mary",
		Age:   12,
		Skill: "cry",
	}
	flag := p.ToJson()
	if !flag {
		t.Fatalf("写文件执行错误，期望值=%v 实际值=%v",true,flag)
	}
	t.Logf("写文件执行正确")
}
func TestReJson(t *testing.T) {
	if !ReJson(){
		t.Fatalf("执行错误，期望值=%v 实际值=%v",true,ReJson())
	}
	t.Logf("执行正确")
}