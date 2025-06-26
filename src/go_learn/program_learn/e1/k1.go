package main

import (
	"fmt"
	"time"
)

func Receive(ch chan int) {
	time.Sleep(1 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("chan close,receive", v)
				return
			} else {
				fmt.Println("receive: ", v)
			}
		}
	}
}

func Send(ch chan int) {
	for i := 0; i < 13; i++ {
		ch <- i
		fmt.Println("send ", i)
	}
	close(ch)
}
func main() {
	ch := make(chan int, 10)

	go Receive(ch)
	go Send(ch)

	for {
		time.Sleep(1 * time.Second)
	}
}
