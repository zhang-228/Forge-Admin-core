package xerr

import (
	"strconv"
)

type XCode interface {
	Error() string
	Code() int
	Message() string
	Details() []interface{}
}

type Code struct {
	code int    // 错误码
	msg  string // 错误信息
}

type Response struct {
	Code   int    `json:"code"`
	Msg    string `json:"message"`
	Result any    `json:"data,omitempty"`
}

func (c Code) Error() string {
	if len(c.msg) > 0 {
		return c.msg
	}

	return strconv.Itoa(c.code)
}

func (c Code) Code() int {
	return c.code
}

func (c Code) Message() string {
	return c.Error()
}

func New(msg string) Code {
	return Code{code: 00, msg: msg}
}

func Add(code int, msg string) Code {
	return Code{code: code, msg: msg}
}

// NewCustomError 自定义错误
func NewCustomError(code int, msg string) error {
	return Code{code: code, msg: msg}
}

// NewRequestParamError 参数错误
func NewRequestParamError(msg string) error {
	return Code{code: 4000, msg: msg}
}

// NewSystemConfError 系统配置文件错误  注：一般用在system配置文件中未配置的信息在代码中使用到的情况
// 例如： 未配置邮箱信息,代码中使用到了需要依赖邮箱信息的代码,会返回该error
func NewSystemConfError(msg string) error {
	return Code{code: 5000, msg: msg}
}
