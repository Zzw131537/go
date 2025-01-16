/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-14 15:56:23
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))
	go func() {
		defer fmt.Println("子go 程结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go 程正在运行，发送元素 = ", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num=", num)
	}
	fmt.Println("main 结束")
}
