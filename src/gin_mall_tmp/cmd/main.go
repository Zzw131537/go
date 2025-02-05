/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-02-05 16:25:14
 */
package main

import (
	"fmt"
	"mall/conf"
	"mall/route"
)

func main() {
	fmt.Println("hello World!")
	conf.Init()
	r := route.NewRouter()
	_ = r.Run(conf.HttpPort)
}
