/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-03-24 10:14:04
 */
// 客户端

package main

import (
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	// 同步调用
	// client, _ := rpc.DialHTTP("tcp", "localhost:1234") // 创建http 客户端，并创建与 1234的链接
	// var result Result
	// if err := client.Call("Cal.Square", 12, &result); err != nil {
	// 	log.Fatal("Failed to call Cal.Square", err)
	// }
	// log.Printf("%d^2 = %d", result.Num, result.Ans)

	// 异步调用
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result
	asyncCall := client.Go("Cal.Square", 12, &result, nil)
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	<-asyncCall.Done
	log.Printf("%d^2 = %d", result.Num, result.Ans)

}
