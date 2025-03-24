package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

// func (cal *Cal) Square(num int) *Result {
// 	return &Result{
// 		Num: num,
// 		Ans: num * num,
// 	}
// }

func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil
}
func main() {
	// 启动 rpc
	rpc.Register(new(Cal)) // 发布Cal 中满足rpc注册条件的方法
	rpc.HandleHTTP()
	log.Printf("server RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}

}
