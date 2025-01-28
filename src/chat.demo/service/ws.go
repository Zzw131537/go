/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-21 19:14:33
 */
package service

import (
	"chat/cache"
	"chat/pkg/e"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const month = 60 * 60 * 24 * 30 // 一个月30天

type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}

type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}

type Boradcast struct {
	Client  *Client
	Message []byte
	Type    int
}

type ClientManager struct {
	Clients    map[string]*Client
	Boradcast  chan *Boradcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client), // 参与连接的用户
	Boradcast:  make(chan *Boradcast),
	Reply:      make(chan *Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

func CreateID(uid, toUid string) string {
	return uid + "->" + toUid // 1->2
}

func Handler(c *gin.Context) {
	uid := c.Query("uid")     // 自己的id
	toUid := c.Query("toUid") // 对方的id
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin解决跨域问题
			return true
		}}).Upgrade(c.Writer, c.Request, nil) // 升级成ws协议
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// 创建一个用户实例
	client := &Client{
		ID:     CreateID(uid, toUid),
		SendID: CreateID(toUid, uid),
		Socket: conn,
		Send:   make(chan []byte),
	}
	// 用户注册到用户管理上
	Manager.Register <- client
	go client.Read()
	go client.Write()
}

func (c *Client) Read() {
	// 用户读操作

	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		c.Socket.PingHandler()
		SendMsg := new(SendMsg)
		//c.Socket.ReadMessage()
		err := c.Socket.ReadJSON(&SendMsg)
		if err != nil {
			fmt.Println("数组格式错误", err)
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		if SendMsg.Type == 1 { // 发送消息
			r1, _ := cache.RedisClient.Get(c.ID).Result() // 1 -> 2
			r2, _ := cache.RedisClient.Get(c.ID).Result() // 2->1
			if r1 > "3" && r2 == "" {                     // 1给2 发3条但2没有回 or 没有看到 就停止 1 发送
				replyMsg := ReplyMsg{
					Code:    e.WebsocketLimit,
					Content: "达到限制",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			} else {
				cache.RedisClient.Incr(c.ID)
				_, _ = cache.RedisClient.Expire(c.ID, time.Hour*24*30*3).Result()
				// 防止过快"分手"
			}
			Manager.Boradcast <- &Boradcast{
				Client:  c,
				Message: []byte(SendMsg.Content), // 发送过来的消息进行广播

			}

		}

	}

}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			replyMsg := ReplyMsg{
				Code:    e.WebsocketSuccess,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(replyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)

		}
	}
}
