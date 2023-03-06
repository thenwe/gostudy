package user

type User struct {
	Name     string `json:"name" binding:"required" form:"name"`
	Password string `json:"password" form:"password"`
	Id       string
}

func (u *User) TableName() string {
	return "socket_users"
}
