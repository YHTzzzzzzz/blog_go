package v1

import (
	"blog_go/handler"
	"blog_go/middleware"
	"blog_go/models/request"
	"github.com/gin-gonic/gin"
)

func RegisterTestRoute(r *gin.Engine) {
	v := r.Group("/v1/test")
	{
		v.GET("ping", handler.PingHandler)
		v.POST("login", middleware.ValidateRequest(&request.TestLoginRequest{}), handler.TestLoginHandler)
	}
}
