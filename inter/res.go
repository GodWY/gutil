package inter

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
)

type Responser interface {
	RetSuccess(detail string, data proto.Message) interface{}
	RetError(err error) interface{}
	RetFail(code int32, detail string) interface{}
}

var responseHandler Responser = &DefaultResp{}

func GetRespHandler() Responser {
	return responseHandler
}

type Result struct {
	Code   int32         `json:"code" `
	Detail string        `json:"detail"`
	Data   proto.Message `json:"data"`
}

type Any struct {
	Result
}

type DefaultResp struct{}

// RetSuccess return success
func (dt *DefaultResp) RetSuccess(detail string, data proto.Message) interface{} {
	return Result{
		Code:   0,
		Detail: detail,
		Data:   data,
	}
}

// RetError error 返回错误
func (dt *DefaultResp) RetError(err error) interface{} {
	r := Result{}
	msg := err.Error()
	_ = json.Unmarshal([]byte(msg), &r)
	return r
}

// RetFail 返回错误结构
func (dt *DefaultResp) RetFail(code int32, detail string) interface{} {
	return Result{
		Code:   code,
		Detail: detail,
		Data:   nil,
	}
}
