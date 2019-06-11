package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
	"webchat/app/service"
	"github.com/gorilla/websocket"
	"log"
	"gopkg.in/fatih/set.v0"
	"sync"
	"encoding/json"
	)

type Node struct {
	Conn *websocket.Conn
	// 并行请求转串行请求
	DataQueue chan []byte
	GroupSets set.Interface
}

var ClientMap = make(map[int]*Node, 0)
var UserService service.UserService
var ContactService service.ContactService
var RwLock sync.RWMutex

const (
	CMD_SINGLE_MSG = 10
	CMD_ROOM_MSG   = 11
	CMD_HEART      = 0
)
type Message struct {
	Id      int  `json:"id,omitempty" form:"id"` //消息ID
	UserId  int  `json:"user_id,omitempty" form:"user_id"` //谁发的
	Cmd     int    `json:"cmd,omitempty" form:"cmd"` //群聊还是私聊
	AddId   int  `json:"add_id,omitempty" form:"add_id"`//对端用户ID/群ID
	Media   int    `json:"media,omitempty" form:"media"` //消息按照什么样式展示
	Content string `json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `json:"pic,omitempty" form:"pic"` //预览图片
	Url     string `json:"url,omitempty" form:"url"` //服务的URL
	Memo    string `json:"memo,omitempty" form:"memo"` //简单描述
	Amount  int    `json:"amount,omitempty" form:"amount"` //其他和数字相关的
}

func Chat(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	// 先鉴权
	query := r.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId,_ := strconv.Atoi(id)
	checkPass:= UserService.CheckToken(userId, token)
	// 然后将http请求升级为WebSocket请求
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return checkPass
		},
	}).Upgrade(w,r,nil)
	if err != nil {
		log.Println(err.Error())
	}
	node := &Node{
		Conn:conn,
		DataQueue:make(chan []byte,50),
		GroupSets:set.New(set.ThreadSafe), // 线程安全 set 是方便一系列的集合操作
	}
	// 查询用户所有的群
	groupIds := ContactService.GetGroupIds(userId)
	if len(groupIds) > 0 {
		for _,groupId := range groupIds {
			node.GroupSets.Add(groupId) // 将所有的群信息加载到set结构中
		}
	}
	RwLock.Lock() // 读写锁
	ClientMap[userId] = node // map关系映射
	RwLock.Unlock() // 读写锁

	// 开始处理发送逻辑
	go sendLogic(node)
	// 处理接受逻辑
	go receiveLogic(node)

}

func sendLogic(node *Node)  {
	for {
		select {
		case data:= <- node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage,data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

func receiveLogic(node *Node)  {
	for {
		_,data,err := node.Conn.ReadMessage()
		log.Println(data)
		if err!=nil{
			log.Println(err.Error())
			return
		}
		go dispatch(data)
		log.Printf("[ws]<=%s\n",data)
	}
}

func dispatch(data []byte)  {
	// 将客户端消息解析到消息体中
	msg := Message{}
	err := json.Unmarshal(data,&msg)
	if err!=nil{
		log.Println(err.Error())
		return
	}
	switch msg.Cmd {
	case CMD_SINGLE_MSG: // 单对单推送的消息
		sendMsg(msg.AddId, data)
	case CMD_ROOM_MSG:
		for _,item := range ClientMap {
			if item.GroupSets.Has(msg.AddId) {
				item.DataQueue <- data
			}
		}
	case CMD_HEART:
	}
}

func sendMsg(userId int,data []byte) {
	RwLock.Lock()
	node,ok := ClientMap[userId]
	RwLock.Unlock()
	if ok {
		node.DataQueue <- data
	}
}

