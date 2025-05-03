package models

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gopkg.in/fatih/set.v0"
)

type Message struct {
	Model
	FormId   int64  `json:"userId"`   // 信息发送者
	TargetId int64  `json:"targetId"` // 信息接收者
	Type     int    // 聊天类型 : 群聊，私聊，广播
	Media    int    // 信息类型 : 文字，图片，音频
	Content  string // 消息内容
	Pic      string `json:"url"` // 图片链接
	Url      string // 文件相关
	Desc     string // 文件描述
	Amount   int    // 其他数据大小
}

func (m *Message) MsgTableName() string {
	return "message"
}

// Node 构造连接
type Node struct {
	Conn      *websocket.Conn // socket 连接
	Addr      string          // 客户端地址
	DataQueue chan []byte     // 消息内容
	GroupSets set.Interface   // 好友 /群
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// Chat 需要: 发送者Id ，接受者Id,消息类型，发送的内容，发送类型
func Chat(w http.ResponseWriter, r *http.Request) {
	//1. 获取参数信息
	query := r.URL.Query()
	Id := query.Get("userId")
	userId, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		zap.S().Info("类型转换失败", err)
		return
	}

	// 升级为socket
	var isvalida = true
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 获取socket 连接,构造消息节点
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	// 将userId 和 Node 绑定
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	// 服务器发送消息
	go sendProc(node)

	// 服务器接收消息
	go recProc(node)
}

// sendProc 从node中获取消息并写入websocket中
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				zap.S().Info("写入消息失败", err)
				return
			}
			fmt.Println("数据发送socket成功")
		}

	}
}

// // recProc 从websocket中将消息拿出，然后解析，再进行消息类型判读，最后将消息发送到目的用户的node 中
// func recProc(node *Node) {
// 	for {
// 		// 获取信息
// 		_, data, err := node.Conn.ReadMessage()
// 		if err != nil {
// 			zap.S().Info("读取消息失败", err)
// 			return
// 		}

// 		// 进行简单实现
// 		msg := Message{}
// 		err = json.Unmarshal(data, &msg)
// 		if err != nil {
// 			zap.S().Info("json解析失败", err)
// 			return
// 		}

// 		if msg.Type == 1 {
// 			zap.S().Info("这是一条私信:", msg.Content)
// 			tarNode, ok := clientMap[msg.TargetId]
// 			if !ok {
// 				zap.S().Info("不存在对应的node", msg.TargetId)
// 				return
// 			}
// 			tarNode.DataQueue <- data
// 			fmt.Println("发送成功: ", string(data))
// 		}
// 	}

// }

// 升级 recProc ，进行udp 连接
func recProc(node *Node) {
	for {
		// 获取信息
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Info("读取信息失败", err)
			return
		}
		// 将消息放进全局channel中，提高系统的并发能力
		brodMsg(data)
	}
}

// 全局channel
var upSendChan chan []byte = make(chan []byte, 1024)

func brodMsg(data []byte) {
	upSendChan <- data
}

// init 方法,运行message包前调用
func init() {
	go UdpSendProc()
	go UdpRecProc()
}

func UdpSendProc() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		// 192.168.31.147
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
		Zone: "",
	})

	if err != nil {
		zap.S().Info("拨号udp端口失败", err)
		return
	}

	defer udpConn.Close()

	for {
		select {
		case data := <-upSendChan:
			_, err := udpConn.Write(data)
			if err != nil {
				zap.S().Info("写入udp消息失败", err)
				return
			}
		}
	}
}

// 完成消息的接收
func UdpRecProc() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})

	if err != nil {
		zap.S().Info("监听udp端口失败", err)
		return
	}

	defer udpConn.Close()

	for {
		var buf [1024]byte
		n, err := udpConn.Read(buf[0:])
		if err != nil {
			zap.S().Info("读取udp数据失败", err)
			return
		}

		// 处理发送逻辑
		dispatch(buf[0:n])
	}
}

// dispatch 解析消息，判断聊天类型
func dispatch(data []byte) {
	// 解析消息
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		zap.S().Info("消息解析失败", err)
		return
	}

	// 判断消息类型
	switch msg.Type {
	case 1: // 私聊
		sendMsg(msg.TargetId, data)
	case 2: // 群发
		sendGroupMsg(uint(msg.FormId), uint(msg.TargetId), data)
	}
}

// sendMsg 向用户单聊发送消息
func sendMsg(id int64, msg []byte) {
	rwLocker.Lock()
	node, ok := clientMap[id]
	rwLocker.Unlock()

	if !ok {
		zap.S().Info("userId 没有对应的node")
		return
	}
	zap.S().Info("targetId: ", id, "node: ", node)
	if ok {
		node.DataQueue <- msg
	}
}

// sendGroupMsg 群发逻辑
func sendGroupMsg(formId, target uint, data []byte) (int, error) {

}
