package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("HUamne Eat....")
}

type SuperMan struct {
	Human // 继承Human

	leval int
}

func main() {

}
