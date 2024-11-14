package router

import (
	v1 "blog_go/pkg/router/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	// 创建 Gin 实例
	r := gin.Default()

	v1.RegisterTestRoute(r)
	v1.RegisterAuthRoutes(r)

	return r
}
