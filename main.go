package main

import (
	"fmt"
	"yingyi.cn/go-universal/crypt"
	"yingyi.cn/go-universal/db"
	"yingyi.cn/go-universal/email"
	"yingyi.cn/go-universal/http"
	"yingyi.cn/go-universal/rcache"
	"yingyi.cn/go-universal/websockets"
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
