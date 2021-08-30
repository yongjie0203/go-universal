package websockets

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/micro/go-micro/util/log"
)

func OnClientMessage(c *Client, data []byte) {
	var message = new(Message)
	err := json.Unmarshal(data, message)
	if err != nil {
		uuid, uerr := uuid.NewUUID()
		if uerr != nil {
			log.Errorf("new uuid error: %s", uerr)
		}
		var head = MessageHead{}
		head.ClientId = c.ClientId
		head.UID = c.UID
		head.TraceId = uuid.String()
		head.AppName = "server"
		head.MessageType = "un_format_msg"
		head.BusinessType = "sys_un_format_msg"
		var body = MessageBody{data}
		message = &Message{Head: head, Body: body}

	}
	OnReceiveMessage(c, message)
}

type Handler func(c *Client, message *Message)

var Handlers = make(map[string]Handler)

func OnReceiveMessage(c *Client, message *Message) {
	handler, ok := Handlers[message.Head.BusinessType]
	if ok {
		handler(c, message)
		return
	}
	UndefinedHandler(c, message)
}

func UndefinedHandler(c *Client, message *Message) {
	uuid, uerr := uuid.NewUUID()
	if uerr != nil {
		log.Errorf("new uuid error: %s", uerr)
	}
	log.Error("handler not fund for " + message.Head.BusinessType)
	var head = MessageHead{}
	head.ClientId = c.ClientId
	head.UID = c.UID
	head.TraceId = uuid.String()
	head.AppName = "server"
	head.MessageType = "handler_not_fund"
	head.BusinessType = "handler_not_fund"
	head.ResponseCode = "404"
	head.ResponseMsg = "handler not fund for " + message.Head.BusinessType
	var body = MessageBody{nil}
	message = &Message{Head: head, Body: body}
	c.SendBusinessMessage(message)
}
