package handler

import (
	"github.com/wanrun-develop/wanrun/internal/auth/adapters/repository"
)

type IAuthHandler interface {
	SignUp() error
	LogIn() error
	LogOut() error
}

type authHandler struct {
	ar repository.IAuthRepository
}

func NewAuthHandler(ar repository.IAuthRepository) IAuthHandler {
	return &authHandler{ar}
}

func (ah *authHandler) SignUp() error { return nil }
func (ah *authHandler) LogIn() error  { return nil }
func (ah *authHandler) LogOut() error { return nil }
