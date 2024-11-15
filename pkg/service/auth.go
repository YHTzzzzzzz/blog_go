package service

import "blog_go/pkg/validator"

type AuthService struct {
}

func (s *AuthService) Login(r *validator.LoginRequest) error {
	return nil
}
