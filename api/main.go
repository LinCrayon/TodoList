package api

import (
	"encoding/json"
	"fmt"
	"todo_list/pkg/e"
	"todo_list/serializer"
)

// ErrorResponse 返回错误信息
func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 40001,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Status: e.SUCCESS,
		Msg:    "成功",
	}
}
