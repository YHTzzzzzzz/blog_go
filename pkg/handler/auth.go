package handler

import (
	"blog_go/pkg/service"
	"blog_go/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req validator.LoginRequest

	// 请求参数绑定结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AuthService.Login(&req); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   "1234",
		"message": "success",
	})
}
