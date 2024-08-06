package router

import (
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/dog/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/dog/controller"
	"github.com/wanrun-develop/wanrun/internal/dog/core/handler"
	"gorm.io/gorm"
)

func NewRouter(e *echo.Echo, dbConn *gorm.DB) {
	dogRepository := repository.NewDogRepository(dbConn)
	dogHandler := handler.NewDogHandler(dogRepository)
	dogController := controller.NewDogController(dogHandler)

	e.GET("/all-dogs", dogController.GetAllDogs)
	e.GET("/dog/:dogID", dogController.GetDogByID)
	e.POST("/create-dog", dogController.CreateDog)
	e.DELETE("/delete-dog", dogController.DeleteDog)
	// e.PUT("/dog/:dogID", dogController.UpateDog)
}
