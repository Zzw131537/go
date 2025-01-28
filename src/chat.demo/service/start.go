/*
 * @Author: Zhouzw
 * @LastEditTime: 2025-01-28 17:10:26
 */
package service

import (
	"chat/conf"
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
			fmt.Printf("有新连接: %s \n", conn.ID)
			fmt.Println("Manager Register", conn)
			Manager.Clients[conn.ID] = conn
			replyMsg := &ReplyMsg{
				Code:    e.WebsocketSuccess,
				Content: "已经连接到服务器了",
			}
			msg, _ := json.Marshal(replyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister:
			fmt.Printf("连接失败! %s", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				replyMasg := &ReplyMsg{
					Code:    e.WebsocketEnd,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(replyMasg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}

		case broadcast := <-Manager.Boradcast:
			message := broadcast.Message
			sendId := broadcast.Client.SendID
			flag := false
			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message:
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			id := broadcast.Client.ID
			if flag {
				replyMsg := &ReplyMsg{
					Code:    e.WebsocketOnlineReply,
					Content: "对方在线应答",
				}
				msg, _ := json.Marshal(replyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err := InsertMsg(conf.MongoDBName, id, string(message), 1, int64(3*month)) // 1 已经读了

				if err != nil {
					fmt.Println("InsertOne Err", err)
				}

			}

		}
	}
}
