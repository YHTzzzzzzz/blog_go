package service

import (
	"blog_go/models/request"
	"fmt"
)

type AuthService struct {
}

func (s *AuthService) Login(r *request.LoginRequest) error {
	fmt.Println(fmt.Sprintf("username: %s, password: %s", r.Username, r.Password))
	return nil
}
