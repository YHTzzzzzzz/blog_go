package router

import (
	"blog_go/middleware"
	v12 "blog_go/router/v1"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	// 创建 Gin 实例
	r := gin.Default()

	// 使用 Recovery()
	r.Use(gin.Recovery())

	// gin 的参数校验 使用全局中间件
	r.Use(middleware.CustomValidation())

	v12.RegisterTestRoute(r)
	v12.RegisterAuthRoutes(r)

	return r
}
