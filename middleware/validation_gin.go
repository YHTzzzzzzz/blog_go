package middleware

import (
	"blog_go/global"
	"blog_go/pkg/utils"
	"blog_go/types/constants"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
)

// 初始化 validator 校验规则
func init() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Fatal("无法获取 validator 实例")
	}

	// 注册自定义校验规则
	if err := v.RegisterValidation("alpha_underscore", alphanumericUnderscore); err != nil {
		log.Fatalf("注册校验规则失败: %v", err)
	}
}

// CustomValidation 自定义错误转换中间件
func CustomValidation() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 执行后续中间件
		c.Next()

		// 处理错误 c.Errors 需要显示的 c.Error(err) 存入，但是会警告 unhandled error 强迫症不建议
		if v, exists := c.Get(constants.ErrorKeyInContextWithValidation); exists {
			// 处理 validator 校验错误
			message := make(map[string]string)
			var validationErrors validator.ValidationErrors
			if errors.As(v.(error), &validationErrors) {
				for _, vErr := range validationErrors {
					message[vErr.Field()] = vErr.Translate(global.TranslatorInstance)
				}
			} else {
				message["system error"] = v.(error).Error()
			}

			var builder strings.Builder
			for _, v := range message {
				builder.WriteString(v + ";")
			}
			// 统一返回错误
			utils.JSONError(c, constants.ResponseCodeParamError, errors.New(builder.String()))
			c.Abort()
		}
	}
}
