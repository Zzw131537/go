/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-14 15:13:43
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.Defer")
		func() {
			defer fmt.Println("B.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
