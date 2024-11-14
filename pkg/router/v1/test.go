package v1

import (
	"blog_go/pkg/handler"
	"github.com/gin-gonic/gin"
)

func RegisterTestRoute(r *gin.Engine) {
	v := r.Group("/v1/test")
	{
		v.GET("ping", handler.PingHandler)
	}
}
