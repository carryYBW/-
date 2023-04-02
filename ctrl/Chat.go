package ctrl

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

const (
	CMD_SINGLE_MSG = 10
	CMD_ROOM_MSG   = 11
	CMD_HEART      = 0
	CMD_ACK        = 1
	CMD_ENTRY_ROOM = 2
	CMD_EXIT_ROOM  = 3
)

type Message struct {
	Id      int64  `json:"id,omitempty" form:"id"`           //消息ID
	Userid  int64  `json:"userid,omitempty" form:"userid"`   //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Dstid   int64  `json:"dstid,omitempty" form:"dstid"`     //对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"`     //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"`         //预览图片
	Url     string `json:"url,omitempty" form:"url"`         //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"`       //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"`   //其他和数字相关的
}

var clientMap map[int64]*Node = make(map[int64]*Node, 0)
var rwlocker sync.RWMutex

func Chat(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	isvalida := checkToken(userId, token)

	//使用 websocket 协议     对原有http协议进行升级
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()

	go sendproc(node)
	go recvproc(node)

	//sendMsg(userId, []byte("hello,world"))
}

// 发送协程
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}

		}
	}

}

// 接收线程
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		dispatch(data)
		fmt.Printf("recv<=%s", data)
	}

}

// todo 参数处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	switch msg.Cmd {
	case CMD_SINGLE_MSG:
		sendMsg(msg.Dstid, data)
	case CMD_ROOM_MSG:
	//群聊转发逻辑
	case CMD_HEART:
		//一般什么都不做  检测心跳

	}

}

// 发送消息
func sendMsg(userId int64, msg []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}

}

func checkToken(userId int64, token string) bool {
	user := userService.Find(userId)
	return user.Token == token
}
