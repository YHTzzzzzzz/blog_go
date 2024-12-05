package types

type ResponseCodeType int

// Response is a struct that represents the response from the API
type Response struct {
	Code    ResponseCodeType `json:"code"`           // 状态码
	Message string           `json:"message"`        // 消息说明
	Data    interface{}      `json:"data,omitempty"` // 返回的数据，可以是任意类型
}

const (
	ResponseCodeSuccess      ResponseCodeType = 200 // 成功
	ResponseCodeError        ResponseCodeType = 500 // 失败
	ResponseCodeTokenInvalid ResponseCodeType = 401 // token无效
	ResponseCodeParamError   ResponseCodeType = 400 // 参数错误
)
