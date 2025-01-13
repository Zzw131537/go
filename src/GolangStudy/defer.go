package main

import "fmt"

func f2() int {
	fmt.Println("call defer")
	return 0
}

func f3() int {
	fmt.Println("call return ")
	return 0
}
func f1() int {

	defer f2()
	return f3()
}

func main() {
	//
	f1()
}
