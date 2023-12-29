package api

import (
	"encoding/json"
	"todo_list_v2.01/pkg/ctl"
	"todo_list_v2.01/pkg/e"
)

// 返回错误信息
func ErrorResponse(err error) *ctl.TrackedErrorResponse {
	//断言变量 err 的类型是否为 *json.UnmarshalTypeError
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(err, "JSON类型不匹配")
	}

	return ctl.RespError(err, "参数错误", e.InvalidParams)
}
