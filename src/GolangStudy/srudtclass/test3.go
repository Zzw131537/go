/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-14 15:10:10
 */
package main

import (
	"fmt"
	"time"
)

/*
协程

groutine
切换从而降低了CPU 的利用率
*/

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine i= %d\n", i)
		time.Sleep(1 * time.Second)
	}

}
