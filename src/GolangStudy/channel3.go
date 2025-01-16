/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-14 16:01:41
 */
package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println("data = ", data)
		} else {
			break
		}
	}
	fmt.Println("Main Finish!")
}
