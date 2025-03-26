package utils

import (
	"blog_go/types"
	"blog_go/types/constants"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JSONResponse 统一 JSON 响应
func JSONResponse(c *gin.Context, code constants.ResponseCodeType, message string, data interface{}, err string) {
	if message == "" {
		message = constants.CodeToMessage[code]
	}
	c.JSON(http.StatusOK, types.Response{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   err,
	})
}

// JSONSuccess 成功返回
func JSONSuccess(c *gin.Context, message string, data interface{}) {
	JSONResponse(c, constants.ResponseCodeSuccess, message, data, constants.DefaultEmpty)
}

// JSONSuccessIgnoreMessage 忽略 message 的成功返回
func JSONSuccessIgnoreMessage(c *gin.Context, data interface{}) {
	JSONResponse(c, constants.ResponseCodeSuccess, "success", data, constants.DefaultEmpty)
}

// JSONError 失败返回
func JSONError(c *gin.Context, code constants.ResponseCodeType, err error) {
	if err == nil {
		err = errors.New("inner system error, please contact the administrator")
	}
	JSONResponse(c, code, constants.DefaultEmpty, nil, err.Error())
}
