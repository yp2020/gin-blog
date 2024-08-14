package model

import (
	"Gin-Blog/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPasswd,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))

	if err != nil {
		fmt.Println("数据库连接失败，请检查参数:", err)
	}
	//表名禁用复数
	db.SingularTable(true)

	//db.AutoMigrate(&User{}, &Category{}, &Article{})
	// 设置连接池中的最大闲置连接数
	db.DB().SetMaxIdleConns(10)
	// 设置数据库中的最大连接数
	db.DB().SetMaxOpenConns(100)
	// 设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(100 * time.Second)
}
