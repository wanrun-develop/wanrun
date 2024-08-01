package router

import (
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/dog/controller"
)

func NewRouter(dc controller.IDogController) *echo.Echo {
	e := echo.New()

	e.GET("/all-dogs", dc.GetAllDogs)
	e.GET("/dog/:dogID", dc.GetDogByID)
	return e
}
