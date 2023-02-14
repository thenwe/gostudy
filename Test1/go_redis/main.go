package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"

)
var(
	ctx=context.Background()
	r *redis.Client
)
type user struct{
	name string
	age int
}
func main(){
	rdb:=redis.NewClient(&redis.Options{Addr:"localhost:6379",Password:"",DB:0})
	
	//set 普通set
	err:=rdb.Set(ctx,"name","Go",0).Err()
	if err!=nil{
		fmt.Println("err=",err)
	}
	val,err:=rdb.Get(ctx,"name").Result()
	if err!=nil{
		fmt.Println("err=",err)
	}
	fmt.Println(val)
	//setnx 不存在才赋值
	err=rdb.SetNX(ctx,"city","mianyang",0).Err()
	if err!=nil{
		fmt.Println("err=",err)
	}
	val1:=rdb.Get(ctx,"city")
	fmt.Println(val1.Val())
	//批量赋值
	sli:=[]string{"zhangsan","lisi"}
	err=rdb.MSet(ctx,"val1",123,"val2",sli[0]).Err()
	if err!=nil{
		fmt.Println("err=",err)
	}
	val2:=rdb.Get(ctx,"val2")
	fmt.Println(val2.Val())
	//原子性批量赋值，要么同时成功，要么同时失败
	err=rdb.MSetNX(ctx,"val2","new233","val3","第三个值").Err()
	if err!=nil{
		fmt.Println("err=",err)
	}
	fmt.Println(rdb.Get(ctx,"val2").Val(),rdb.Get(ctx,"val3").Val())
	//列表从左插入
	i:=rdb.LLen(ctx,"list").Val()//获取列表长度
	if i<8{
		//列表从左插入
		err=rdb.LPush(ctx,"list","21",123,23,2323232).Err()
		if err!=nil{
			fmt.Println("err=",err)
		}
	}
	fmt.Println("列表数据：",rdb.LRange(ctx,"list",0,-1).Val())
	//下标获取元素
	_,err=rdb.LIndex(ctx,"list",2).Result()
	if err!=nil{
		fmt.Println("err=",err)
	}
	fmt.Println("列表第三个元素:",rdb.LIndex(ctx,"list",2).Val())
	//判断列表是否存在
	j,err:=rdb.Exists(ctx,"list").Result()
	if err!=nil{
		fmt.Println("err=",err)
	}
	if j==1{
		fmt.Println("list存在")
	}else {
		fmt.Println("list不存在")
	}
	//从左移除list
	if rdb.LLen(ctx,"list").Val()>8{

		rdb.LPop(ctx,"list")//移除最左边的元素
		fmt.Println("移除最左边元素后的列表数据：",rdb.LRange(ctx,"list",0,-1).Val())
		//从右移除list
		_,err=rdb.RPop(ctx,"list").Result()//移除最右边的元素
		fmt.Println("移除最右边元素后的列表数据：",rdb.LRange(ctx,"list",0,-1).Val())
		//指定元素移除
		k,err:=rdb.LRem(ctx,"list",1,21).Result()//指定移除一个"21"
		if err!=nil{
			fmt.Println("err=",err)
			return
		}
		if k==1{
			fmt.Println("移除21成功，移除后的列表数据为",rdb.LRange(ctx,"list",0,-1).Val())
		}else{
			fmt.Println("移除失败")
		}
		//截取下标范围的元素
		_,err=rdb.LTrim(ctx,"list",1,5).Result()//截取下标1到5的元素
		if err!=nil{
			fmt.Println("err=",err)
			return
		}
		fmt.Println("截取成功，截取后的list数据为:",rdb.LRange(ctx,"list",0,-1).Val())
		//移除列表最后一个元素并移动到新列表中 （暂时有问题）
		/*if rdb.LLen(ctx,"list").Val()>6{
			_,err=rdb.RPopLPush(ctx,"list","list2").Result()
			if err!=nil{
				fmt.Println("err=",err)
				return
			}
			fmt.Println("移动成功，移动后的原list为:",rdb.LRange(ctx,"list",0,-1).Val())
			fmt.Println("移动成功，移动后的新list为:",rdb.LRange(ctx,"list2",0,-1).Val())
		}*/
	}
	//在列表指定元素前后插入值
	_,err=rdb.LInsert(ctx,"list","before","23","新茶入职左").Result()
	if err!=nil{
		fmt.Println("err=",err)
		return
	}
	_,err=rdb.LInsert(ctx,"list","after","23","新茶入职右").Result()
	if err!=nil{
		fmt.Println("err=",err)
		return
	}
	fmt.Println("插入成功，插入后的list数据为:",rdb.LRange(ctx,"list",0,-1).Val())
	//列表指定下标赋值
	_,err=rdb.LSet(ctx,"list",1,"下标为1的元素赋值").Result()
	if err!=nil{
		fmt.Println("err=",err)
		return
	}
	fmt.Println("赋值成功，赋值后的list数据为:",rdb.LRange(ctx,"list",0,-1).Val())
	//列表指定下标获取元素
	v:=rdb.LIndex(ctx,"list",1).Val()
	fmt.Println("获取成功，list下标为1数据为:",v)



}