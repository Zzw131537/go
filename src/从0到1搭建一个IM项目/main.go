package main

import (
	"HiChat/initialize"
	"HiChat/router"
)

func main() {
	// 初始化日志
	initialize.InitLogger()
	//初始化配置
	initialize.InitConfig()
	//初始化数据库
	initialize.InitDB()
	initialize.InitRedis()

	router := router.Router()
	router.Run(":8000")
}
