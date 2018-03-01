package proto

import (
	"encoding/json"
	"log"
)

type Code int

var (
	Ok                  Proto
	InvalidDataFormat   Proto
	InvalidParameter    Proto
	ServiceNotAvailable Proto
)

func RegisterOk(p Proto) {
	Ok = p
}

func RegisterInvalidDataFormat(p Proto) {
	InvalidDataFormat = p
}

func RegisterInvalidParameter(p Proto) {
	InvalidParameter = p
}

func RegisterServiceNotAvailable(p Proto) {
	ServiceNotAvailable = p
}

type Proto interface {
	Code() int
	Text() string
}

type Reply struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func NewReply(result interface{}) *Reply {
	return newReply(Ok, result)
}

func ReplyCode(p Proto) *Reply {
	return newReply(p, nil)
}

func newReply(p Proto, result interface{}) *Reply {
	return &Reply{
		Code:    p.Code(),
		Message: p.Text(),
		Result:  result,
	}
}

func (r *Reply) Json() []byte {
	data, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
		r = ReplyCode(ServiceNotAvailable)
		data, _ = json.Marshal(r)
	}
	return data
}
