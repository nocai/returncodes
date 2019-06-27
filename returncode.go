package returncodes

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

// 状态码对应着httpStatus，业务异常请定义 >= 1000(4位数)
var (
	Success     = New(http.StatusOK, "操作成功")
	ErrSystem   = NewErrorCoder(http.StatusInternalServerError, "系统异常")
	ErrArgument = NewErrorCoder(http.StatusBadRequest, "参数有误")
	ErrTimeout  = NewErrorCoder(http.StatusRequestTimeout, "处理超时")
)

type ReturnCoder interface {
	Code() int
	Message() string
	Data() interface{}
}

type returnCode struct {
	C int         `json:"Code"`
	M string      `json:"Message"`
	D interface{} `json:"Data,omitempty"`
}

func (rc returnCode) Code() int {
	return rc.C
}

func (rc returnCode) Message() string {
	return rc.M
}

func (rc returnCode) Data() interface{} {
	return rc.D
}

var _ ReturnCoder = returnCode{}

type ErrorCoder interface {
	ReturnCoder
	error
}

type errorCode struct {
	returnCode
}

func (ec errorCode) Error() string {
	return ec.Message()
}

var _ ErrorCoder = errorCode{}

func New(code int, message string) ReturnCoder {
	return &returnCode{C: code, M: message}
}

func NewErrorCoder(code int, message string) ErrorCoder {
	return &errorCode{
		returnCode: returnCode{C: code, M: message},
	}
}

func Fail(i interface{}) ErrorCoder {
	switch i.(type) {
	case error:
		err := errors.Cause(i.(error))
		if err, ok := err.(ErrorCoder); ok {
			return err
		}
		return NewErrorCoder(ErrSystem.Code(), err.Error())
	default:
		return NewErrorCoder(ErrSystem.Code(), fmt.Sprint(i))
	}
}

func Succ(message string, data interface{}) ReturnCoder {
	if message == "" && data == nil {
		panic(errors.New("invalid args: message and data are all zero value"))
	}
	return &returnCode{C: Success.Code(), M: message, D: data}
}

func Mess(message string) ReturnCoder {
	return Succ(message, nil)
}

func Data(data interface{}) ReturnCoder {
	return Succ(Success.Message(), data)
}
