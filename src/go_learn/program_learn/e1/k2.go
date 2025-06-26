package main

import (
	"fmt"
	"sync"
	"time"
)

// 实现并发安全就需要加锁
type Money struct {
	lock   sync.Mutex
	amount int64
}

func (m *Money) Add(x int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.amount += x
}

func (m *Money) Minute(x int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.amount -= x
}

func (m *Money) Get() int64 {
	return m.amount
}

func main() {
	m := new(Money)
	m.Add(1000)
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(50 * time.Millisecond)
			m.Minute(5)
		}()
	}
	time.Sleep(10 * time.Second)
	fmt.Println(m.Get())
}
