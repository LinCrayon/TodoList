package types

// CreateTaskReq 创建备忘录请求
type CreateTaskReq struct {
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

// ListTaskReq 分页查询用户
type ListTaskReq struct {
	Limit int `json:"limit" form:"limit"` //一页多少个
	Start int `json:"start" form:"start"` //第几页
}

type ListTaskResp struct {
	Id        int64  `json:"id"`      // 任务ID
	Title     string `json:"title"`   // 题目
	Content   string `json:"content"` // 内容
	View      int64  `json:"view"`    // 浏览量
	Status    int    `json:"status"`  // 状态(0未完成，1已完成)
	CreatedAt int64  `json:"created_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

type ShowTaskReq struct {
	Id int64 `json:"id" form:"id"`
}

type UpdateTaskReq struct {
	Id      int64  `form:"id" json:"id"`
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content" `
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

type SearchTaskReq struct {
	Info string `form:"info" json:"info"`
}

type DeleteTaskReq struct {
	Id int64 `json:"id" form:"id"`
}
