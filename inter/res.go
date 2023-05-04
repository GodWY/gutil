package internal

import (
	"encoding/json"
	"github.com/gogo/protobuf/proto"
)

type Responser interface {
	RetSuccess(detail string, data proto.Message) Result
	RetError(err error) Result
	RetFail(code int32, detail string) Result
}

var responseHandler Responser = &DefaultResp{}

func GetRespHandler() Responser {
	return responseHandler
}

type Result struct {
	Code   int32         `json:"code,omitempty" `
	Detail string        `json:"detail,omitempty"`
	Data   proto.Message `json:"data"`
}

type DefaultResp struct{}

// RetSuccess return success
func (dt *DefaultResp) RetSuccess(detail string, data proto.Message) Result {
	return Result{
		Code:   0,
		Detail: detail,
		Data:   data,
	}
}

// RetError error 返回错误
func (dt *DefaultResp) RetError(err error) Result {
	r := Result{}
	msg := err.Error()
	_ = json.Unmarshal([]byte(msg), &r)
	return r
}

// RetFail 返回错误结构
func (dt *DefaultResp) RetFail(code int32, detail string) Result {
	return Result{
		Code:   code,
		Detail: detail,
		Data:   nil,
	}
}
