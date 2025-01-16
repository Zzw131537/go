/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-14 15:42:18
 */
package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine 结束")
		fmt.Println("groutine 正在运行")
		c <- 666

	}()
	num := <-c
	fmt.Println("num = ", num)
	fmt.Println("maingoroutine 结束")
}
