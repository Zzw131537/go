package main

import (
	"fmt"
	"sync"
)

// 栈
type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.array = append(stack.array, v)
	stack.size += 1
}

func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty")
	}

	v := stack.array[stack.size-1]

	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray

	stack.size -= 1
	return v
}

func (stack *ArrayStack) Top() string {
	if stack.size == 0 {
		panic("empty")
	}
	v := stack.array[stack.size-1]
	return v
}

func (stack *ArrayStack) Size() int {
	return stack.size
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
}
