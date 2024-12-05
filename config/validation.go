package config

import (
	"blog_go/global"
	"blog_go/types/constants"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// InitValidator 初始化 validator 翻译器 global.Translator
func InitValidator() {
	// 修改gin框架中的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册获取json tag的方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get(constants.LanguageZh), ",", 2)[0]
			if name == "-" || name == "" {
				name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			}
			return name
		})

		// 创建翻译器实例
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT)

		// 根据当前配置获取对应语言
		locale := global.ServerConfigInstance.AppConfigInstance.Locale
		global.TranslatorInstance, ok = uni.GetTranslator(locale)
		if !ok {
			fmt.Println("Error loading translator for locale:", locale)
		}
		if global.TranslatorInstance == nil {
			fmt.Println("Error loading translator for locale:", locale)
			global.TranslatorInstance, _ = uni.GetTranslator(constants.LanguageZh) // 默认使用中文
		}

		// 注册翻译器
		var err error
		switch locale {
		case constants.LanguageEn:
			err = enTranslations.RegisterDefaultTranslations(v, global.TranslatorInstance)
		case constants.LanguageZh:
			err = zhTranslations.RegisterDefaultTranslations(v, global.TranslatorInstance)
		default:
			err = zhTranslations.RegisterDefaultTranslations(v, global.TranslatorInstance)
		}

		if err != nil {
			fmt.Println("Error registering translations:", err)
		}
	}
}
