package service

import (
	user "Test1/gin1/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

const Url = "root:236332@tcp(162.14.64.254:3306)/first_test?charset=utf8mb4&loc=Local"

var Db, err = gorm.Open(mysql.Open(Url), &gorm.Config{})

func ToJson(c *gin.Context) {
	dto := user.User{}
	//name := c.DefaultQuery("name", "woziji")
	err := c.ShouldBind(&dto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": "0",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  dto,
		"code": "1",
	})
}
func ToRegister(c *gin.Context) {
	dto := user.User{}
	var dto1 []user.User
	//name := c.DefaultQuery("name", "woziji")
	err := c.ShouldBind(&dto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": "0",
		})
		return
	}
	err = Db.Debug().Where(dto).Find(&dto1).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": "0",
		})
		return
	}
	if len(dto1) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户存在",
			"code": "1",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户不存在",
			"code": "1",
		})
		return
	}
}
func ToUpdate(c *gin.Context) {
	dto := user.User{}
	var dto1 []user.User
	//name := c.DefaultQuery("name", "woziji")
	err := c.ShouldBind(&dto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": "0",
		})
		return
	}
	Db.Debug().Where(dto).Find(&dto1)

	err = Db.Debug().Model(&dto1).Update("name", "测试名字").Error//根据主键更新
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  err.Error(),
			"code": "1",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "更新成功",
			"code": "1",
		})
		return
	}
}
