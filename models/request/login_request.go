package request

type LoginRequest struct {
	Username string `json:"username" binding:"required" zh:"用户名"`
	Password string `json:"password" binding:"required" zh:"密码"`
}
