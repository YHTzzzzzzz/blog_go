package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
)

// 创建全局的 validator 实例
var validate = validator.New()

// 自定义校验函数：仅允许大小写字母、数字和下划线 _
func alphanumericUnderscore(fl validator.FieldLevel) bool {
	regex := `^[a-zA-Z0-9_]+$`
	return regexp.MustCompile(regex).MatchString(fl.Field().String())
}

/*
*
在 Go 语言中，init() 函数会在 包被导入时自动执行，因此不需要显式调用 init()。
为什么 init() 不需要手动调用？
init() 是 Go 语言的特殊函数，会在 包初始化时自动执行一次。
只要你的 middleware 包 被 import 了，init() 就会运行，无需手动调用。
*/
func init() {
	// 注册自定义校验规则
	err := validate.RegisterValidation("custom", alphanumericUnderscore)
	if err != nil {
		panic(errors.New("register validation failed"))
	}
}

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
