/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-15 19:05:08
 */
package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net dail error:", err)
		return nil
	}
	client.conn = conn
	return client

}

func (client *Client) menu() bool {
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	var flag int

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("输入不合法")
		return false
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}

		// 处理不同业务
		switch client.flag {
		case 1:
			fmt.Println("公聊模式选择...")
			break

		case 2:
			fmt.Println("私聊模式选择....")
			break

		case 3:
			fmt.Println("更新用户名选择....")
			break
		}
	}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器ip地址")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口")
}
func main() {

	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>连接失败>>>>>")
		return
	}

	fmt.Println(">>>>>>连接成功>>>>>")

	client.Run()
}
