package model

import (
	"Gin-Blog/utils/errmsg"
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:" type:varchar(20);not null" json:"username"`
	Password string `gorm:" type:varchar(100);not null" json:"password"`
	Role     int    `gorm:" type:int" json:"role"`
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// CheckUser 查询用户是否存在
func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	fmt.Println("user.ID", user.ID)
	// 用户已经存在
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// GetUsers  查询用户列表
func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteUsers 删除用户
func DeleteUsers(id int) int {
	err = db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 修改密码

// BeforeSave  用钩子函数来实现加密
//func (u *User) BeforeSave() {
//	u.Password = ScryptPw(u.Password)
//}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	salt := []byte{123, 1, 2, 45, 65, 46, 76, 98}
	dk, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	finalPassword := base64.StdEncoding.EncodeToString(dk)
	return finalPassword
}
