/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-20 17:58:32
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
