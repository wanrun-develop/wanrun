package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/wanrun-develop/wanrun/internal/dog/core/handler"
	"github.com/wanrun-develop/wanrun/pkg/errors"
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
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve dog information",
		})
	}

	return c.JSON(http.StatusOK, resDogs)
}

func (dc *dogController) GetDogByID(c echo.Context) error {
	dogIDStr := c.Param("dogID")
	dogID, err := strconv.Atoi(dogIDStr)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid dog ID format",
		})
	}
	resDog, err := dc.dh.GetDogByID(uint(dogID))
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve dog information",
		})
	}

	return c.JSON(http.StatusOK, resDog)
}
