package middleware

import (
	"blog_go/pkg/config"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CustomValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 直接执行 其他中间件 && 逻辑处理
		c.Next()

		// check 错误
		if len(c.Errors) > 0 {
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
				}
				errMessages = append(errMessages, message)
			}
			// 如果需要，可以对错误进一步处理，例如自定义返回信息
			c.JSON(http.StatusInternalServerError, gin.H{"errors": errMessages})
		}
	}
}

func Error(err error) (message string) {
	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		return err.Error()
	} else {
		for _, e := range validationErrors {
			message += e.Translate(config.Trans) + ";"
		}
	}
	return message
}
