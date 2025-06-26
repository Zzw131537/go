package main

import (
	"fmt"
	"reflect"
)

type A interface {
	Println()
}

type B interface {
	Println()
	Printf() int
}

type A1Instance struct {
	Data string
}

func (a1 *A1Instance) Println() {
	fmt.Println("a1: ", a1.Data)
}

type A2Instance struct {
	Data string
}

func (a2 *A2Instance) Println() {
	fmt.Println("a2 :", a2.Data)
}

func (a2 *A2Instance) Printf() int {
	fmt.Println("a2 :", a2.Data)
	return 0
}

func main() {

	var a A
	a = &A1Instance{Data: "132"}
	a.Println()

	if v, ok := a.(*A1Instance); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A1")
	}

	fmt.Println(reflect.TypeOf(a).String())

	a = &A1Instance{Data: "5555"}

	a.Println()
	if v, ok := a.(*A2Instance); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not a A2")
	}
	fmt.Println(reflect.TypeOf(a).String())

	var b B
	b = &A2Instance{Data: "999"}
	fmt.Println(b.Printf())
}
