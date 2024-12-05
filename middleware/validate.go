package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// 创建全局的 validator 实例
var validate = validator.New()

// ValidateRequest 中间件：自动验证请求体的结构体 todo 优化：错误提示信息
func ValidateRequest(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求中绑定 JSON 数据到结构体
		if err := c.ShouldBindJSON(&model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "请求参数格式错误"})
			c.Abort()
			return
		}

		// 验证结构体
		if err := validate.Struct(model); err != nil {
			// 使用类型断言时检查是否为 ValidationErrors 类型
			var validationErrs validator.ValidationErrors
			if errors.As(err, &validationErrs) {
				var errorMessages map[string]string
				// 遍历所有的验证错误信息，提取字段名和对应的错误信息
				errorMessages = make(map[string]string)
				for _, e := range validationErrs {
					// 将字段名和错误信息添加到结果中
					errorMessages[e.Field()] = e.Tag()
				}
				c.JSON(http.StatusBadRequest, gin.H{"error": errorMessages})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.Abort()
			return
		}

		// 请求通过验证，继续处理
		c.Next()
	}
}

//
