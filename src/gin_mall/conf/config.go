/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 15:05:37
 */
package conf

import (
	"mall/dao"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppModel string
	HttpPort string

	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDb         string
	RedisAddrstring string
	RedisPw         string
	RedisDbName     string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	// 本地读取环境变量
	file, err := ini.Load("C:/Users/86131/go/src/gin_mall/conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMySql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPhotpPath(file)

	// mysql 读 主
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")

	// mysql 写 从
	PathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")

	dao.Database(pathRead, PathWrite)

}

func LoadServer(file *ini.File) {
	AppModel = file.Section("service").Key("AppModel").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMySql(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddrstring = file.Section("redis").Key("RedisAddrstring").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("Email").Key("ValidEmail").String()
	SmtpHost = file.Section("Email").Key("SmtpHost").String()
	SmtpEmail = file.Section("Email").Key("SmtpEmail").String()
	SmtpPass = file.Section("Email").Key("SmtpPass").String()
}

func LoadPhotpPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}
