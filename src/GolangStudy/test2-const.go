/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-07 23:40:23
 */
package main

import (
	"fmt"
)

const (
	// 可以在const () 添加iota ，每行iota 累加1 ,第一行iota 为0
	BEIJIING = 10 * iota
	SHANGHAI
	SHENZHENG
)

const (
	a, b = iota + 1, iota + 2 // a = 1 b = 2
	c, d                      // 2 , 3
	e, f                      // 3,4

	g, h = iota * 2, iota * 3 // iota = 3  g=6,h=9
	i, k                      // iota  i 8,12
)

func main() {
	const length = 10
	fmt.Println("length = ", length)

	fmt.Println(BEIJIING, SHANGHAI, SHENZHENG)
	fmt.Println(g, h)
	fmt.Println(i, k)

}
