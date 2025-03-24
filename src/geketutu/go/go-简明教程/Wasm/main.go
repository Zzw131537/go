/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-24 10:34:47
 */
package main

import "syscall/js"

func main() {
	//第一步，新建文件 main.go，使用 js.Global().get(‘alert’) 获取全局的 alert 对象，通过 Invoke 方法调用。等价于在 js 中调用 window.alert("Hello World")。
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World!")
}
