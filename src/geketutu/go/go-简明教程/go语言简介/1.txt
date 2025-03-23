/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-21 15:35:52
 */
package main

import (
	"fmt"
	// "internal/stringslite"
	// "reflect"
)

type Student struct {
	age  int
	name string
}

func (stu *Student) hello(name string) string {

	return fmt.Sprintf("Hello %s,I am %s", name, stu.name)
}
func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func main() {
	fmt.Println("Hello World")

	// str1 := "Golang"
	// str2 := "Go 语言"

	// fmt.Println(reflect.TypeOf(str2[2]).Kind()) // uint8
	// fmt.Println(str1[2], string(str1[2]))
	// fmt.Printf("%d %c \n", str2[2], str2[2])
	// fmt.Println("len(str2) :", len(str2))

	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		arr[i] += 100
	}
	fmt.Println(arr)

	// m1 := make(map[string] int)

	// m2 := map[string] string {
	// 	"Sam" : "Male",
	// 	"Alice":"Female",
	// }

	// m1["Tome"] = 18
	stu := &Student{
		name: "Tome",
		age:  18,
	}
	msg := stu.hello("Jack")
	fmt.Println(msg)

}
