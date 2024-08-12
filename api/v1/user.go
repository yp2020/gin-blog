package v1

import (
	"Gin-Blog/model"
	"Gin-Blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserExit 查询用户是否存在
func UserExit(c *gin.Context) {

}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "参数绑定错误",
		})
		return
	}

	code := model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   user,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 查询单个用户。本博客没有社交功能，这个功能意义不大

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	users := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   users,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

}
