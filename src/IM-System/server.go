/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-15 18:35:34
 */
package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播channel
	Message chan string
}

// 创建一个server 的接口
func NewServer(ip string, port int) *Server {

	server := Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return &server
}

func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg

		}
		this.mapLock.Unlock()
	}
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + " : " + msg

	this.Message <- sendMsg
}
func (this *Server) Handler(conn net.Conn) {
	// .. 当前链接的业务
	//	fmt.Println("链接建立成功")
	// 用户上线，加入OnlineMpa 然后广播
	user := NewUser(conn, this)

	user.Online()

	isLive := make(chan bool)
	// 广播上线消息

	// 结束客户端传送的数据
	go func() {
		buf := make([]byte, 4000)
		for {
			n, err := conn.Read(buf)

			if n == 0 {
				user.OffLine()

				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read error: ", err)
				return
			}

			msg := (string)(buf[:n-1])

			// 进行广播
			user.DoMessage(msg)

			isLive <- true

		}
	}()

	// 阻塞
	for {

		select {
		case <-isLive:

		case <-time.After(time.Second * 300):
			// 已经超时
			user.SendMeg("你被踢了")

			close(user.C)
			conn.Close()

			return
		}

	}

}

// 启动服务器的方法
func (this *Server) Start() {
	// sockrt listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Println("net.error: ", err)
		return
	}

	// close listen socket
	defer listener.Close()

	// 监听msg

	go this.ListenMessage()

	for {
		// accept

		conn, er := listener.Accept()

		if er != nil {
			fmt.Println("listener.accept error : ", er)
			continue
		}

		// do handler
		go this.Handler(conn)
	}

}
