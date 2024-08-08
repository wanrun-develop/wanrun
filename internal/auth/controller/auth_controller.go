package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/auth/core/handler"
)

type IAuthController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

type authController struct {
	ah handler.IAuthHandler
}

func NewAuthController(ah handler.IAuthHandler) IAuthController {
	return &authController{ah}
}

func (ac *authController) SignUp(c echo.Context) error { return nil }
func (ac *authController) LogIn(c echo.Context) error  { return nil }
func (ac *authController) LogOut(c echo.Context) error { return nil }
