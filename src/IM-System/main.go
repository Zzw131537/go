/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-15 14:26:53
 */
package main

import "time"

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
	time.Sleep(10 * time.Second)
}
