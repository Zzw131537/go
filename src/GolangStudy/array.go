/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-13 15:23:45
 */
package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println("len = %d ,slice = %v\n", len(a), a)

	var myMap map[string]string
	if myMap == nil {
		fmt.Println("map1 is null")
	}
	myMap = make(map[string]string, 10) // 使用前需要分配空间

	myMap2 := make(map[string]string)
	myMap2["one"] = "1"

}
