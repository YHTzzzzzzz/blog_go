package utils

import (
	"blog_go/types"
	"blog_go/types/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func JSONSuccess(c *gin.Context, message string, data interface{}) {
	JSONResponse(c, constants.ResponseCodeSuccess, message, data, "")
}

func JSONSuccessIgnoreMessage(c *gin.Context, data interface{}) {
	JSONResponse(c, constants.ResponseCodeSuccess, constants.DefaultEmpty, data, "")
}

func JSONError(c *gin.Context, code constants.ResponseCodeType, message string, err error) {
	// 处理 validator 翻译
	var errMessage string
	if code == constants.ResponseCodeParamError {
		errMessage = ValidateErrorTranslator(err)
	} else {
		errMessage = err.Error()
	}
	JSONResponse(c, code, message, nil, errMessage)
}
