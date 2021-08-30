package websockets

import "time"

type hub struct {
	clients map[string]*Client

	connect    chan *Client // 连接处理
	disConnect chan *Client
	broadcast  chan []byte //广播消息内容
}

var Hub = hub{
	clients:    make(map[string]*Client),
	connect:    make(chan *Client),
	disConnect: make(chan *Client),
	broadcast:  make(chan []byte),
}

func StartHub() {
	for {
		select {
		case c := <-Hub.connect:
			Hub.clients[c.ClientId] = c
			go c.StartClient()
			go c.startRead()
		case d := <-Hub.disConnect:
			d.Conn.Close()
			delete(Hub.clients, d.ClientId)
			d.IsDeleted = true
		case b := <-Hub.broadcast:
			go Broadcast(b)
		default:
			time.Sleep(time.Microsecond * 200)
		}
	}

}

func Broadcast(data []byte) {
	for _, c := range Hub.clients {
		c.Write <- data
	}
}
