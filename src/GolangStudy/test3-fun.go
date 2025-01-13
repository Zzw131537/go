/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-11 18:34:16
 */
// 多返回值
package main

import "fmt"

func fool(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)
	c := 100
	return c
}

func fool2(a string, b int) (int, int) {
	fmt.Println(a)
	fmt.Println(b)
	return 666, 777
}
func main() {

	c := fool("anc", 555)
	fmt.Println(c)
	res1, res2 := fool2("bbb", 444)
	fmt.Println(res1)
	fmt.Println(res2)
}
