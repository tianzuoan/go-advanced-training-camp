package geekerror

import (
	"fmt"
	"week02/error/errorcode"
)

var CodeNotFound Interface

func init() {
	CodeNotFound = NewGeekError(errorcode.NOT_FOUND)
}

type Interface interface {
	Error() string
	SetUserMsg(format string, a ...interface{}) Interface
	GetUserMsg() string
	GetCode() int
}

// NewGeekError 新建GeekError
func NewGeekError(code int) Interface {
	geekErr := &GeekError{
		Code: code,
		Msg:  errorcode.GetErrMsg(code),
	}
	return geekErr
}

// GeekError geek error struct
type GeekError struct {
	Code int
	Msg  string
}

func (e *GeekError) Error() string {
	return e.Msg
}

func (e *GeekError) SetUserMsg(format string, a ...interface{}) Interface {
	e.Msg = fmt.Sprintf(format, a...)
	return e
}

func (e *GeekError) GetUserMsg() string {
	return e.Msg
}

func (e *GeekError) GetCode() int {
	return e.Code
}
