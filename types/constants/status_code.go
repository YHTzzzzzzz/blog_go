package constants

type ResponseCodeType int

const (
	ResponseCodeSuccess       ResponseCodeType = 200 // 成功
	ResponseCodeInternalError ResponseCodeType = 500 // 服务器内部错误
	ResponseCodeTokenInvalid  ResponseCodeType = 401 // token无效
	ResponseCodeParamError    ResponseCodeType = 400 // 参数错误
)

var CodeToMessage = map[ResponseCodeType]string{
	ResponseCodeSuccess:       "成功",
	ResponseCodeInternalError: "服务器内部错误",
	ResponseCodeTokenInvalid:  "token无效",
	ResponseCodeParamError:    "参数错误",
}
