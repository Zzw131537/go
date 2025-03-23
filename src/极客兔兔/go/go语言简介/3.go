/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-21 15:45:28
 */
package main

import (
	"fmt"
	"time"
)

var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道

func download(url string) {
	fmt.Println("start to download ", url)
	time.Sleep(time.Second)
	ch <- url
}
func main() {
	for i := 0; i < 3; i++ {
		go download("a.com/" + string(i+'0'))

	}
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println("finish ", msg)

	}
	fmt.Println("Down!")
}
