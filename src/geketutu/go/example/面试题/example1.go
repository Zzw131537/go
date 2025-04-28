package main

/*
实现一个批处理任务，list=[1,2,3,4,5,6,7,8]，对list中的元素进行平方运算，要求goroutine数量最多3个。
*/
import (
	"fmt"
	"sync"
)

func power(n int) int {
	return n * n
}

type ParalleRunner struct {
	wg     sync.WaitGroup
	sendQ  chan int
	result map[int]int
	mu     sync.Mutex
}

func NewParalleRunner(parallel int) *ParalleRunner {
	r := &ParalleRunner{
		sendQ:  make(chan int),
		result: make(map[int]int),
	}

	for i := 0; i < parallel; i++ {
		r.wg.Add(1)
		go func() {
			defer r.wg.Done()

			for v := range r.sendQ {
				r.mu.Lock()
				r.result[v] = power(v)
				r.mu.Unlock()
			}
		}()
	}
	return r
}

func (r *ParalleRunner) Run(num int) {
	r.sendQ <- num
}

func (r *ParalleRunner) Wait() {
	close(r.sendQ)
	r.wg.Wait()
}

func (r *ParalleRunner) Result() map[int]int {
	return r.result
}

func main() {
	r := NewParalleRunner(3)
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, v := range list {
		r.Run(v)
	}
	r.Wait()
	dict := r.Result()
	var res []int
	for _, v := range list {
		res = append(res, dict[v])
	}
	fmt.Println(res)
}
