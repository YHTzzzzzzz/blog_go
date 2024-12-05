package middleware

import (
	"blog_go/config"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

//	todo 重复错误 结构{
//	   "error": "Key: 'LoginRequest.username' Error:Field validation for 'username' failed on the 'required' tag"
//	}{
//
//	   "errors": [
//	       "username为必填字段;"
//	   ]
//	}
//
// CustomValidation 自定义错误转换器 -- 需要显示的使用 c.Error(err) 添加错误
func CustomValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 直接执行 其他中间件 && 逻辑处理
		c.Next()

		// check 错误
		if len(c.Errors) > 0 { // 这里使用 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) 无法记录到
			// 自定义错误提示
			var errMessages []string
			var validationErrors validator.ValidationErrors
			for _, e := range c.Errors {
				// check 是否是 validator 错误
				message := ""
				if errors.As(e, &validationErrors) {
					for _, e := range validationErrors {
						message += e.Translate(config.Trans) + ";"
					}
				} else {
					message = e.Error()
				}
				errMessages = append(errMessages, message)
			}
			// 如果需要，可以对错误进一步处理，例如自定义返回信息
			c.JSON(http.StatusInternalServerError, gin.H{"errors": errMessages})
		}
	}
}
