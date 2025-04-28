package main

import (
	"fmt"
	"sync"
)

// 使用两个goroutine，交替打印数字和字母

func main() {
	digit := make(chan struct{})
	letter := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		startIndex := 1
		for {
			if _, ok := <-digit; !ok {
				break
			}
			if startIndex >= 10 {
				close(letter)
				break
			}

			fmt.Printf("%d", startIndex)
			startIndex++
			letter <- struct{}{}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		startIndex := 'a'
		for {
			if _, ok := <-letter; !ok {
				break
			}
			fmt.Printf("%c", startIndex)
			startIndex++
			digit <- struct{}{}
		}
	}()
	digit <- struct{}{}
	wg.Wait()
}
