package db

import (
	"context"
	"fmt"
	"gin/dto"
	url2 "gin/url"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var url=url2.Url
var Db,err=gorm.Open(mysql.Open(url),&gorm.Config{})
var rdb=redis.NewClient(&redis.Options{Addr:"localhost:6379",Password:"",DB:0})
var ctx=context.Background()

func Socket(user *dto.SocketUser) bool{
	if err!=nil{
		fmt.Println("连接失败:",err)
		return false
	}
	err=Db.AutoMigrate(&dto.SocketUser{})
	if err != nil{
		fmt.Println("迁移失败:",err)
		return false
	}
	result:=Db.Create(user)
	if result.Error!=nil{
		println("创建记录失败:",err)
		return false
	}
	fmt.Println("创建成功")
	return true
}

func Register(c *gin.Context) {
	var sUser []dto.SocketUser
	name:=c.Query("account")
	result:=Db.Debug().Where("name=?",name).Find(&sUser)
	//SELECT * FROM `socket_users` WHERE name='admin' 用对象接收时返回一条记录 用切片接收时可返回多条记录
	if result.Error!=nil{
		fmt.Println("查询失败")
		c.JSON(200,gin.H{
			"flag":false,
			"message":"查询失败",
		})
		return
	}
	fmt.Println(sUser)
	if result.RowsAffected>=1{
		c.JSON(200,gin.H{
			"flag":false,
			"message":"用户已存在",
		})
		return
	}
	//创建用户
	user1:=dto.SocketUser{Name:c.Query("account"),Password:c.Query("password")}
	result=Db.Create(user1)
	if result.Error!=nil{
		c.JSON(200,gin.H{
			"flag":false,
			"message":result.Error,
		})
		return
	}
	c.JSON(200,gin.H{
		"flag":true,
		"message":"注册成功",
	})
}

//中间件
func BeforeLogin(c *gin.Context){
	name:=c.Query("account")
	var sUser []dto.SocketUser
	//查询用户是否已经存在
	result:=Db.Debug().Where("name=?",name).Find(&sUser)
	if result.RowsAffected==0{
		c.JSON(200,gin.H{
			"flag":false,
			"message":"用户不存在",
		})
		//fmt.Println(c.Request)
		c.Abort()//让后面的函数不再执行
		return
	}
	c.Next()
}

//登录
func Login(c *gin.Context){
	//用户名存redis
	_,err:=rdb.Set(ctx,"user",c.Query("account"),0).Result()
	if err!=nil{
		fmt.Println("存redis失败:",err)
		c.JSON(200,gin.H{
			"flag":false,
			"message":"存redis失败",
		})
		return
	}
	//登录时默认将renwei添加为好友
	rdb.SAdd(ctx,c.Query("account"),"renwei")
	var fUser1 []dto.UserFriend
	var fUser dto.UserFriend
	fUser.Name=c.Query("account")
	fUser.FriendName="renwei"
	Db.Debug().Where(fUser).Find(&fUser1)
	//SELECT * FROM user_friends WHERE name = "admin" AND friend_name = "renwei"
	if len(fUser1)<1{
		Db.Debug().Create(fUser)
	}
	fmt.Println("---------->")
	c.JSON(200,gin.H{
		"flag":true,
		"message":"登陆成功",
	})
}