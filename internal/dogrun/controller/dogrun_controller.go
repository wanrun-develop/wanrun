package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/wanrun-develop/wanrun/internal/dogrun/core/handler"
	"github.com/wanrun-develop/wanrun/pkg/errors"
)

type IDogrunController interface {
	GetDogrunDetail(echo.Context) error
}

type DogrunController struct {
	h handler.IDogrunHandler
}

func NewDogrunController(h handler.IDogrunHandler) IDogrunController {
	return &DogrunController{h}
}

func (dc *DogrunController) GetDogrunDetail(c echo.Context) error {
	placeId := c.Param("placeId")

	dogrun, err := dc.h.GetDogrunDetail(placeId)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve dog run information",
		})
	}

	return c.JSON(http.StatusOK, dogrun)
}
