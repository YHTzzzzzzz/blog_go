package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func TestLoginHandler(c *gin.Context) {
	//var req request.TestLoginRequest
	//
	//// 请求参数绑定结构体
	//if err := c.ShouldBindJSON(&req); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	//c.Error(err) 这里是给 c.Errors() 塞内容
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{
		"token":   "1234",
		"message": "success",
	})
}
