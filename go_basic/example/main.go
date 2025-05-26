package main

import "fmt"

// 泛型

type MySlice[T int | float64] []T

func (s *MySlice[T]) Sum() T {
	var sum T
	for _, v := range *s {
		sum += v
	}
	return sum
}
func main() {

	// 泛型函数
	var s MySlice[int] = []int{1, 2, 3}
	fmt.Println(s.Sum())
}
