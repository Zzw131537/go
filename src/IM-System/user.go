/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-15 18:34:08
 */
package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户的api
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	go user.ListenMessage()

	return user
}

// 用户上线业务
func (this *User) Online() {

	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播上线消息
	this.server.BroadCast(this, "已上线")
}

// 用户下线业务
func (this *User) OffLine() {

	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)

	this.server.mapLock.Unlock()

	// 广播上线消息
	this.server.BroadCast(this, "已下线")

}

// 给当前用户对应的客户端发送消息
func (this *User) SendMeg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlingMsg := "[" + user.Addr + "]" + user.Name + ": " + "在线.....\n"
			this.SendMeg(onlingMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]

		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMeg("当前用户已存在")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMeg("你的用户名已更新")
		}

	} else if len(msg) > 4 && msg[0:3] == "to|" {

		remotname := strings.Split(msg, "|")[1]

		if remotname == "" {
			this.SendMeg("消息格式错误")
			return
		}

		remotUser, ok := this.server.OnlineMap[remotname]
		if !ok {
			this.SendMeg("用户名不存在")
			return
		}
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMeg("消息不存在，请重发")
			return
		}
		remotUser.SendMeg(this.Name + "对你说:" + content)

	} else {
		this.server.BroadCast(this, msg)
	}

}

// 监听当前User channel 的方法
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}
