package main

import "fmt"

func main() {
	m := make(map[string]int64, 4)

	m["dog"] = 1
	m["cat"] = 1
	m["hen"] = 1
	fmt.Println(m)

	which := "hen"
	v, ok := m[which]
	if ok {
		fmt.Println("找到了")
	} else {
		fmt.Println("not find:", which)
	}
	which = "ccc"
	v, ok = m[which]
	if ok {
		// 找到了
		fmt.Println("find:", which, "value:", v)
	} else {
		// 找不到
		fmt.Println("not find:", which)
	}
}
