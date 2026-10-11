package code

// RespCode 响应码类型
type RespCode int

const (
	Fail         RespCode = -1  // 失败
	Success      RespCode = 0   // 成功
	Unauthorized RespCode = 401 // 未授权
	ServerError  RespCode = 500 // 服务器错误
)
