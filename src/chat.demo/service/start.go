/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-21 18:46:21
 */
package service

import (
	"chat/pkg/e"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func (manage *ClientManager) Start() {
	for {
		fmt.Println("----------监听管道通信-----------")
		select {
		case conn := <-Manager.Register:
			fmt.Printf("有新连接: %v", conn.ID)
			Manager.Clients[conn.ID] = conn
			replyMsg := ReplyMsg{
				Code:    e.WebsocketSuccess,
				Content: "已经连接到服务器了",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)

		}
	}
}
