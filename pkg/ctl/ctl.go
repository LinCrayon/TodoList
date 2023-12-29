package ctl

import (
	"todo_list_v2.01/pkg/e"
)

// Response 最基础的返回resp
type Response struct {
	Status int    `json:"status"`
	Data   any    `json:"data"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}

// DataList 带有总数的Data结构 ,数组形式返回
type DataList struct {
	Item  any   `json:"item"`
	Total int64 `json:"total"`
}

// TrackedErrorResponse 追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	TrackId string `json:"track_id"`
}

// RespList 带有总数的列表构建器
func RespList(items any, total int64) *Response {
	return &Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}

// RespSuccess 正常返回成功的函数
func RespSuccess(code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   "操作成功",
		Msg:    e.GetMsg(status),
	}
	return r
}

// RespSuccessWithData 带data成功返回
func RespSuccessWithData(data any, code ...int) *Response {
	status := e.SUCCESS
	if code != nil {
		status = code[0]
	}

	r := &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}

	return r
}

// RespError 错误返回
func RespError(err error, data string, code ...int) *TrackedErrorResponse {
	status := e.ERROR
	if code != nil {
		status = code[0]
	}

	r := &TrackedErrorResponse{
		Response: Response{
			Status: status,
			Msg:    e.GetMsg(status),
			Data:   data,
			Error:  err.Error(),
		},
		//TrackId: "", // TODO 加上track id
	}

	return r
}
