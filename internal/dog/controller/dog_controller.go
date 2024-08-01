package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/dog/core/handler"
)

type IDogController interface {
	GetAllDogs(c echo.Context) error
	GetDogByID(c echo.Context) error
}

type dogController struct {
	dh handler.IDogHandler
}

func NewDogController(dh handler.IDogHandler) IDogController {
	return &dogController{dh}
}

func (dc *dogController) GetAllDogs(c echo.Context) error {
	resDogs, err := dc.dh.GetAllDogs()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	return c.JSON(http.StatusOK, resDogs)
}

func (dc *dogController) GetDogByID(c echo.Context) error {
	dogIDStr := c.Param("dogID")
	dogID, _ := strconv.Atoi(dogIDStr)
	resDog, err := dc.dh.GetDogByID(uint(dogID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error)
	}

	return c.JSON(http.StatusOK, resDog)
}
