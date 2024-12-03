package v1

import (
	"blog_go/pkg/handler"
	"blog_go/pkg/service"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	authHandler := handler.AuthHandler{AuthService: &service.AuthService{}}
	authGroup := r.Group("/v1/auth")
	{
		// todo 非显示使用middleware进行校验
		authGroup.POST("login", authHandler.Login)
	}

}
