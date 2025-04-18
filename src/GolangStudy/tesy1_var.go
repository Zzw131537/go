/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-07 23:29:40
 */
// 变量的声明方弄1�7
package main

import (
	"fmt"
)

// 声明全局变量
var gA int = 100
var gB = 200

// 只能在局部使用方泄1�74

func main() {
	// 方法丢�: 默认值为0
	var a int
	fmt.Println("a的��是 ", a)

	// f2: 初始化一个��1�7
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("Type of b = %T\n", b)

	// 方法3 : 省去数据结构自动匹配类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("Type of c = %T\n", c)

	// 方法4: 省去 var
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Printf("type of f = %T\n", f)

	g := 1.16
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("GA = ", gA, " Gb = ", gB)

	// 声明多个变量
	var xx, yy = 100, 200
	fmt.Println(xx, yy)
	var kk, ll = 100, "Abc"
	fmt.Println(kk, ll)

	var (
		v1 int     = 100
		v2 float32 = 300.2
	)        
	fmt.Println(v1, v2)

}
