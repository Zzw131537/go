/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-17 14:34:30
 */
package main

import (
	"chat/conf"
	"chat/router"
)

func main() {
	conf.Init()

	r := router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
