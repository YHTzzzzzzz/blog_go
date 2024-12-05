package handler

import (
	"blog_go/models/request"
	"blog_go/models/response"
	"blog_go/pkg/service"
	"blog_go/pkg/utils"
	"blog_go/types/constants"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	// 请求参数绑定结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONError(c, constants.ResponseCodeParamError, constants.DefaultEmpty, err)
		return
	}

	if token, err := h.AuthService.Login(&req); err != nil {
		utils.JSONError(c, constants.ResponseCodeInternalError, constants.DefaultEmpty, err)
	} else {
		resp := &response.LoginResponse{Token: token}
		utils.JSONSuccess(c, "login success", resp)
	}
}
