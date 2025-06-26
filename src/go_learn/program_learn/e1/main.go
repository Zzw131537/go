package main

import (
	"fmt"
	"time"
)

func HU(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("after 2 second hu !!")
	ch <- 100
}
func main() {

	ch := make(chan int)

	go HU(ch)
	fmt.Println("start hu ,wait...")

	// 读取空信道会阻塞，直到消息到来
	v := <-ch
	fmt.Println("receive: ", v)

}
