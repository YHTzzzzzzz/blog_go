package utils

import (
	"blog_go/global"
	"errors"
	"github.com/go-playground/validator/v10"
)

// ValidateErrorTranslator 参数校验翻译器
func ValidateErrorTranslator(err error) string {
	var message string
	var validateErrors validator.ValidationErrors
	if ok := errors.As(err, &validateErrors); ok {
		for _, e := range validateErrors {
			message += e.Translate(global.TranslatorInstance) + ";"
		}
	} else {
		message = err.Error()
	}
	return message
}
