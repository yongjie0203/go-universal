package websockets

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/util/log"
	"strings"
	"time"
)

type Client struct {
	ClientId    string          // 标识ID
	Conn        *websocket.Conn // 用户连接
	ConnectTime uint64          // 首次连接时间
	IsDeleted   bool            // 是否删除或下线
	UID         int64           // 业务端标识用户ID
	Read        chan []byte
	Write       chan []byte
}

func NewClient(clientId string, socket *websocket.Conn) *Client {
	return &Client{
		ClientId:    clientId,
		Conn:        socket,
		ConnectTime: uint64(time.Now().Unix()),
		IsDeleted:   false,
		Read:        make(chan []byte),
		Write:       make(chan []byte),
	}
}

func (c *Client) StartClient() {

	for {
		select {
		case r := <-c.Read:
			fmt.Printf("收到客户端消息： %s\n", string(r))
			var msg = string(r)
			log.Info("客户端信息：", msg)
			log.Info("len(r):", len(r))
			log.Info("len([]byte(\"0\")):", len([]byte("0")))
			log.Info("r：", r)
			log.Info("[]byte(\"0\") ：", []byte("0"))
			log.Info("是否ping:", strings.EqualFold("0", msg))
			if strings.EqualFold("0", msg) { //响应ping
				log.Info("响应ping")
				var pong = "1"
				go func() { c.Write <- []byte(pong) }()

			} else {
				go OnClientMessage(c, r)
			}
		case w := <-c.Write:
			fmt.Printf("向客户端发送消息： %s\n", string(w))
			err := c.Conn.WriteMessage(websocket.TextMessage, w)

			if err != nil {
				log.Error(err)
			}
		default:
			//time.Sleep(time.Microsecond * 200 )
		}
	}
}

func (c *Client) startRead() {
	go func() {
		for {
			msgType, msg, err := c.Conn.ReadMessage()
			fmt.Printf("msgtype: %s err:%s", msgType, err)
			c.Read <- msg
		}
	}()
}

func (c *Client) SendMessage(data []byte) {
	go func() { c.Write <- data }()
}

func (c *Client) SendJsonMessage(o *interface{}) {
	data, err := json.Marshal(o)
	if err != nil {
		log.Error(err)
	}
	c.SendMessage(data)
}

// SendBusinessMessage 发送业务消息给客户端
func (c *Client) SendBusinessMessage(msg *Message) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Error(err)
	}
	c.SendMessage(data)
}
