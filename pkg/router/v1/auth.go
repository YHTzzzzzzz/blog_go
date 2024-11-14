package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/v1/auth")
	{
		authGroup.POST("login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}
}
