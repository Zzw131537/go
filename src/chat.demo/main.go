/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-21 18:41:41
 */
package main

import (
	"chat/conf"
	"chat/router"
	"chat/service"
)

func main() {
	conf.Init()

	go service.Manager.Start()

	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)

}
