package request

type TestLoginRequest struct {
	Username string `json:"username" validate:"required,custom"`
	Password string `json:"password" validate:"required"`
}
