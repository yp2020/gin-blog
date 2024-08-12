package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

// 读取 config.ini 配置文件
//
//	并赋给对应的值
var (
	AppMode  string
	HttpPort string
	Db       string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPasswd string
	DbName   string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
		panic(err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("debug")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPasswd = file.Section("database").Key("DbPassword").MustString("y123p456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
