/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-21 15:41:10
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download ", url)
	time.Sleep(time.Second)
	wg.Done()
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download("a.com/" + string(i+'0'))

	}
	wg.Wait()
	fmt.Println("Down!")
}
