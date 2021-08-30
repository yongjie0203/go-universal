package http

type HttpReturn struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *HttpReturn {
	return Return(data, 1, "")
}

func Return(data interface{}, code int, message string) *HttpReturn {
	r := new(HttpReturn)
	r.Code = code
	r.Data = data
	r.Message = message
	return r
}
