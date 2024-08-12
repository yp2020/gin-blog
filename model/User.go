package model

import (
	"Gin-Blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:" type:varchar(20);not null" json:"username"`
	Password string `gorm:" type:varchar(20);not null" json:"password"`
	Role     int    `gorm:" type:int" json:"role"`
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	// 用户已经存在
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}
