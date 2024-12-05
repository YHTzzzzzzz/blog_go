package types

import "blog_go/types/constants"

// Response is a struct that represents the response from the API
type Response struct {
	Code    constants.ResponseCodeType `json:"code"`            // 状态码
	Message string                     `json:"message"`         // 消息说明
	Data    interface{}                `json:"data,omitempty"`  // 返回的数据，可以是任意类型
	Error   string                     `json:"error,omitempty"` // 错误信息
}
