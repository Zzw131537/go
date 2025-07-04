package main

import "fmt"

func feib(x int) int {
	if x == 0 {
		return 1
	}
	return x * feib(x-1)
}

// 二分查找

// 二分查找递归解法
func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		// 出界了，找不到
		return -1
	}

	// 从中间开始找
	mid := (l + r) / 2
	middleNum := array[mid]

	if middleNum == target {
		return mid // 找到了
	} else if middleNum > target {
		// 中间的数比目标还大，从左边找
		return BinarySearch(array, target, l, mid-1)
	} else {
		// 中间的数比目标还小，从右边找
		return BinarySearch(array, target, mid+1, r)
	}

}

func main() {
	arr := []int{1, 5, 9, 15, 81, 89, 123, 189, 333}
	target := 500
	res := BinarySearch(arr, target, 0, len(arr)-1)
	fmt.Println(target, res)
	target = 189
	res = BinarySearch(arr, target, 0, len(arr)-1)
	fmt.Println(target, res)
}
