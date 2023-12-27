package serializer

// Response 基础序列化器
type Response struct {
	Status int    `json:"status"`
	Data   any    `json:"data"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
}
