package main

import "fmt"

// 单一职责原则
type Clothes struct {
}

func (c *Clothes) OnWork() {
	fmt.Println("工作的装扮")
}

func (c *Clothes) OnShop() {
	fmt.Println("逛街的装扮")
}

type ClothesWork struct{}

func (cw *ClothesWork) Style() {
	fmt.Println("逛街的装扮")
}

func main() {
	c := Clothes{}

	fmt.Println("在工作")
	c.OnWork()

	c.OnShop()
}
