package middleware

import (
	"blog_go/pkg/config"
	"blog_go/pkg/global"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

// 创建全局的 validator 实例
var validate = validator.New()

// 创建全局的翻译器实例
var trans ut.Translator

// InitValidator 初始化 validator 翻译器 todo 翻译器未生效
func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册获取json tag的方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 创建翻译器实例
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT)

		// 根据当前配置获取对应语言
		locale := config.ServerConfigInstance.AppConfigInstance.Locale
		trans, _ = uni.GetTranslator(locale)
		if trans == nil {
			fmt.Println("Error loading translator for locale:", locale)
			trans, _ = uni.GetTranslator(global.LanguageZh) // 默认使用中文
		}

		// 注册翻译器
		var err error
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		}

		if err != nil {
			fmt.Println("Error registering translations:", err)
		}
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
