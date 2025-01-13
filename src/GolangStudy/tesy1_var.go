/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-07 23:29:40
 */
// 鍙橀噺鐨勫０鏄庢柟寮�
package main

import (
	"fmt"
)

// 澹版槑鍏ㄥ眬鍙橀噺
var gA int = 100
var gB = 200

// 鍙兘鍦ㄥ眬閮ㄤ娇鐢ㄦ柟娉�4

func main() {
	// 鏂规硶涓€: 榛樿鍊间负0
	var a int
	fmt.Println("a鐨勫€兼槸 ", a)

	// f2: 鍒濆鍖栦竴涓€�
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("Type of b = %T\n", b)

	// 鏂规硶3 : 鐪佸幓鏁版嵁缁撴瀯鑷姩鍖归厤绫诲瀷
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("Type of c = %T\n", c)

	// 鏂规硶4: 鐪佸幓 var
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Printf("type of f = %T\n", f)

	g := 1.16
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("GA = ", gA, " Gb = ", gB)

	// 澹版槑澶氫釜鍙橀噺
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
