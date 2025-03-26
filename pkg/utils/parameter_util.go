package utils

import (
	"blog_go/types/constants"
	"github.com/gin-gonic/gin"
)

// BindAndValidate 统一绑定参数并校验错误
func BindAndValidate(c *gin.Context, obj interface{}) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.Set(constants.ErrorKeyInContextWithValidation, err)
		c.Abort()
		return false
	}
	return true
}
