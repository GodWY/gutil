package gutil

import (
	"encoding/json"
	"github.com/gogo/protobuf/proto"
)

type Result struct {
	Code   int32         `json:"code"`
	Detail string        `json:"detail"`
	Data   proto.Message `json:"data"`
}

// RetSuccess return success
func RetSuccess(detail string, data proto.Message) Result {
	return Result{
		Code:   0,
		Detail: detail,
		Data:   data,
	}
}

// RetError error 返回错误
func RetError(err error) Result {
	r := Result{}
	msg := err.Error()
	_ = json.Unmarshal([]byte(msg), &r)
	return r
}

// RetFail 返回错误结构
func RetFail(code int32, detail string) Result {
	return Result{
		Code:   code,
		Detail: detail,
		Data:   nil,
	}
}
