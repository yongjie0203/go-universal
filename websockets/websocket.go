package websockets

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/util/log"

	"net/http"
	"time"
	"yingyi.cn/go-universal/rcache"
)

// 心跳间隔
var heartbeatInterval = 25 * time.Second

// PingTimer 启动定时器进行心跳检测
func PingTimer() {
	go func() {
		ticker := time.NewTicker(heartbeatInterval)
		defer ticker.Stop()
		for {
			<-ticker.C
			//发送心跳
			for clientId, client := range Hub.clients {
				if err := client.Conn.WriteControl(websocket.PingMessage, []byte{1}, time.Now().Add(time.Second)); err != nil {
					Hub.disConnect <- client
					log.Errorf("发送心跳失败: %s 总连接数：%d", clientId, len(Hub.clients))
				}
			}
		}

	}()
}

func SetUpWebSocket() {
	go StartHub()
	ws := gin.New()
	ws.Any("/", func(c *gin.Context) {
		//2.建立websocket连接
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

		if err != nil {
			fmt.Println(err)
		}

		id := rcache.NextId("WEBSOCKET_CLIENT_ID")
		client := NewClient(string(id), conn)
		Hub.connect <- client

	})

	ws.Run(":8081")

}

//1.设置websocket参数，CheckOrigin表示是否允许跨域
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebSocketClientId() string {
	return ""
}
