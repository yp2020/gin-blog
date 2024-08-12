package errmsg

const (
	SUCCESS                = 200
	ERROR                  = 500
	ERROR_USERNAME_USED    = 1001 //code =100x 用户模块的错误
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	// code =200x 文章模块的错误

	// code =3000 分类模块的错误
)

// 错误 code 与 msg 的映射关系
var codeMsg = map[int]string{
	SUCCESS:                "ok",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN 已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

// GetErrMsg 错误处理机制，拿到状态码，查询对应的 msg 返回给前端。
func GetErrMsg(code int) string {
	return codeMsg[code]
}
