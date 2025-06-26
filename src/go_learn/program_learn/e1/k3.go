package main

import "fmt"

// 汉诺塔

var cnt = 0

func main() {
	n := 4
	a := "a"
	b := "b"
	c := "c"
	tower(n, a, b, c)
	fmt.Println(cnt)
}

func tower(n int, a, b, c string) {
	if n == 1 {
		cnt = cnt + 1
		fmt.Println(a, "->", c)
		return
	}
	tower(n-1, a, c, b)
	cnt = cnt + 1
	fmt.Println(a, "->", c)
	tower(n-1, b, a, c)

}
