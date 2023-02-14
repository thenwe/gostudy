package main

import (
	"fmt"
	"gin/db"
	"gin/dto"
)

func main(){
	//fmt.Println(db.QueryBefore("admin"))
	var fUser1 []dto.UserFriend
	var fUser dto.UserFriend
	fUser.Name="admin"
	fUser.FriendName="renwei"
	db.Db.Debug().Where(fUser).Find(&fUser1)
	//SELECT * FROM user_friends WHERE name = "admin" AND friend_name = "renwei"
	if len(fUser1)<1{
		db.Db.Debug().Create(fUser)
		return
	}
	fmt.Println("未查询到")
}