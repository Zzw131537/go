package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex

var set = make(map[int]bool, 0)

func printOne(num int) {
	mu.Lock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
	mu.Unlock()
}

func main() {
	for i := 0; i < 10; i++ {
		go printOne(i)
	}
	time.Sleep(time.Second)
}
