package proto

import (
	"stock/comm/http/proto"
)

type Code int

const (
	// system code
	Ok                  Code = 200
	ServiceNotAvailable Code = iota // 0
	ActionNotFound                  // 1
	InvalidDataFormat               // 2
	InvalidParameter                // 3
	Unauthorized                    // 4
	TokenExpired                    // 5
	Forbidden                       // 6
	TimeOut                         // 7
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 20

	// application common code
	AccountExists           // 21
	AccountNotExists        // 22
	IncorrectPassword       // 23
	InvalidCaptchaCode      //24
	CaptchaApplyFrequently  //25
	SMSLimit                // 26
	MobileAssociated        // 27
	MobileHasBeenAssociated // 28
	AnswerCardNotExists     // 29
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_
	_ // 40

	// application code
	SchoolNotExists // 41
	OrderNotExists  // 42
)

var codeText = map[Code]string{
	Ok:                  "ok",
	ServiceNotAvailable: "当前服务不可用, 请稍后再试",
	ActionNotFound:      "HTTP方法或请求路径有误",
	InvalidDataFormat:   "请求数据格式有误",
	InvalidParameter:    "请求参数有误",
	Unauthorized:        "用户名或密码错误",
	TokenExpired:        "会话已过期，请重新登陆",
	Forbidden:           "您不被允许做此操作",
	TimeOut:             "请求超时",

	AccountExists:           "该账号已存在，请直接登陆",
	AccountNotExists:        "该账号不存在",
	IncorrectPassword:       "密码错误",
	InvalidCaptchaCode:      "验证码无效",
	CaptchaApplyFrequently:  "验证码申请过于频繁，请稍后再试",
	SMSLimit:                "短信下发次数已达上限，详情请联系客服",
	MobileAssociated:        "您已绑定过手机号码",
	MobileHasBeenAssociated: "该手机号码已被绑定",

	SchoolNotExists: "未找到该学校",
	OrderNotExists:  "订单不存在",
}

func (c Code) Code() int {
	return int(c)
}

func (c Code) Text() string {
	return codeText[c]
}

func init() {
	proto.RegisterOk(Ok)
	proto.RegisterInvalidDataFormat(InvalidDataFormat)
	proto.RegisterInvalidParameter(InvalidParameter)
	proto.RegisterServiceNotAvailable(ServiceNotAvailable)
}
