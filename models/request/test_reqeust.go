package request

type TestLoginRequest struct {
	Username string `json:"username" validate:"required,inputCheck" zh:"用户名"`
	Password string `json:"password" validate:"required" zh:"密码"`
}
