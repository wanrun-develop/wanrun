package core

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/dogOwner/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/dogOwner/core/models"
)

/*
dog_ownerのhandler
*/
type dogOwnerHandler struct {
	repo *repository.DogOwnerRepository
}

/*
handlerの作成
*/
func NewDogOwnerHandler(repo *repository.DogOwnerRepository) *dogOwnerHandler {
	return &dogOwnerHandler{repo}
}

/*
dog_ownerの作成
*/
func (h *dogOwnerHandler) CreateDogOwner(c echo.Context) error {

	dogOwner := models.DogOwner{}

	if err := c.Bind(&dogOwner); err != nil {
		return err
	}
	h.repo.InsertDogOwner(&dogOwner)
	return c.JSON(http.StatusCreated, dogOwner)
}
