package conf

import (
	"fmt"
	"strings"
	"user/model"

	ini "gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}
	LoadMysqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").string()
	DbHost = file.Section("mysql").Key("DbHost").string()
	DbPort = file.Section("mysql").Key("DbPort").string()
	DbUser = file.Section("mysql").Key("DbUser").string()
	DbPassWord = file.Section("mysql").Key("DbPassWord").string()
	DbName = file.Section("mysql").Key("DbName").string()
}
