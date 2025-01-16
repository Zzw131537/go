/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-13 16:56:28
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 200, 10, []string{"xingye", "zhongxingci"}}
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	// 将[]byte转换为字符串
	fmt.Printf("json = %s\n", jsonStr) // 或者 fmt.Println(string(jsonStr))
}
