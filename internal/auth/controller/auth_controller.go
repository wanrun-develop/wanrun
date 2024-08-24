package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/wanrun-develop/wanrun/internal/auth/core/handler"
	model "github.com/wanrun-develop/wanrun/internal/models"
	"github.com/wanrun-develop/wanrun/pkg/errors"
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

func (ac *authController) SignUp(c echo.Context) error {
	reqADO := model.ReqAuthDogOwner{}
	if err := c.Bind(&reqADO); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid format",
		})
	}

	// SignUp
	resAuthDogOwner, err := ac.ah.SignUp(&reqADO)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to register dog owner information",
		})
	}

	return c.JSON(http.StatusOK, resAuthDogOwner)
}

func (ac *authController) LogIn(c echo.Context) error  { return nil }
func (ac *authController) LogOut(c echo.Context) error { return nil }
