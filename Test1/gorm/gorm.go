package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type student struct {
	Name      string
	StuNumber string
	Class     string
}

var url = "root:236332@tcp(162.14.64.254:3306)/first_test?charset=utf8mb4&loc=Local"
var db1, err = gorm.Open(mysql.Open(url), &gorm.Config{})

func createMethod(tx gorm.DB) {

}
func main() {
	str1 := "Lucy"
	number1 := "20231116"
	var msg error = nil
	stu := student{
		Name:      str1,
		StuNumber: number1,
		Class:     "三年级二班",
	}
	var stuList []student
	var stuList1 []student
	student1 := append(stuList, stu)

	if err != nil {
		fmt.Printf("连接失败%v\n", err)
	}
	fmt.Println("连接成功")

	db1.Debug().Where("name = ?", stu.Name).Find(&stuList1) //SELECT * FROM `students` WHERE name = '张三'

	db1.Transaction(func(tx *gorm.DB) error {
		tx.Debug().Create([]student{{"Bruce", "123215", "五年级五班"}})
		//创建数据
		if len(stuList1) < 1 {
			er := tx.Debug().Create(student1).Error //INSERT INTO `students` (`name`,`stu_number`,`class`) VALUES ('张三','3120170','三年级二班')
			if er != nil {
				return er
			}
		} else {
			fmt.Println("该用户已存在")
		}
		//更新数据
		updateData := "20231116"
		tx.Debug().Where("name = ?", str1).Find(&stuList1)
		fmt.Println(stuList1)
		if len(stuList1) < 1 {
			fmt.Println("更新用户不存在")
		} else {
			fmt.Println("---------")
			er := tx.Debug().Model(new(student)).Where("name = ?", str1).Update("stu_number", updateData).Error //UPDATE `students` SET `stu_number`='20231116' WHERE name = '李四'
			if er != nil {
				return er
			}
		}
		//更新数据2 更新多列
		stu1 := student{
			Name:      "Tom",
			Class:     "六年级五班",
			StuNumber: "20231117",
		}
		er := tx.Debug().Model(new(student)).Where("name = ?", "张三").Updates(stu1).Error
		if er != nil {
			fmt.Println(er)
			return er
		}

		return msg

	})

}
