package main

import (
	"Gin-Blog/model"
	"Gin-Blog/routes"
)

func main() {
	// 初始化数据库
	model.InitDb()
	routes.InitRouter()
}
