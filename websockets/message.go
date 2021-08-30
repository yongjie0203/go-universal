package websockets

type Message struct {
	Head MessageHead `json:"head"`
	Body MessageBody `json:"body"`
}

type MessageHead struct {
	TraceId      string `json:"trace_id"`
	PrevTraceId  string `json:"prev_trace_id"`
	ResponseCode string `json:"response_code"`
	ResponseMsg  string `json:"response_msg"`
	AppName      string `json:"app_name"`
	AppCode      string `json:"app_code"`
	MessageType  string `json:"message_type"`
	BusinessType string `json:"business_type"`
	Endpoint     string `json:"endpoint"`
	ClientId     string `json:"client_id"`
	UID          int64  `json:"uid"`
}

type MessageBody struct {
	Data interface{} `json:"data"`
}
