package main

import (
	"fmt"
	"github.com/yongjie0203/go-universal/crypt"
	"github.com/yongjie0203/go-universal/db"
	"github.com/yongjie0203/go-universal/email"
	"github.com/yongjie0203/go-universal/http"
	"github.com/yongjie0203/go-universal/rcache"
	"github.com/yongjie0203/go-universal/websockets"
)

func main() {
	fmt.Println("hello")
	crypt.Encode("")
	rcache.SetupRedis()
	db.NextId("ddd")
	email.SendMail([]string{"1@qq.com"}, "", "")
	http.Success("")
	websockets.SetUpWebSocket()
}
