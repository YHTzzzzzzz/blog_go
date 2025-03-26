package request

type TestLoginRequest struct {
	Username string `json:"username" validate:"required,inputCheck"`
	Password string `json:"password" validate:"required"`
}
