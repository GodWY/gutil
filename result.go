package gutil

import (
	"github.com/GodWY/gutil/inter"
	"google.golang.org/protobuf/proto"
	"sync"
)

// RetSuccess return success
func RetSuccess(detail string, data proto.Message) interface{} {
	if responseHandler == nil {
		return inter.GetRespHandler().RetSuccess(detail, data)
	}
	return responseHandler.RetSuccess(detail, data)
}

// RetError error 返回错误
func RetError(err error) interface{} {
	if responseHandler == nil {
		return inter.GetRespHandler().RetError(err)
	}
	return responseHandler.RetError(err)
}

// RetFail 返回错误结构
func RetFail(code int32, detail string) interface{} {
	if responseHandler == nil {
		return inter.GetRespHandler().RetFail(code, detail)
	}
	return responseHandler.RetFail(code, detail)
}

var responseHandler = inter.GetRespHandler()
var mu sync.Mutex

// RegisterHttpResponse 注册错误处理使用
func RegisterHttpResponse(h inter.Responser) {
	mu.Lock()
	defer mu.Lock()
	responseHandler = h
}
