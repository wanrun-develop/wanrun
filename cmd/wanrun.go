package wanruncmd

import (
	"log"

	"github.com/wanrun-develop/wanrun/internal/db"
	"github.com/wanrun-develop/wanrun/internal/dog/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/dog/controller"
	"github.com/wanrun-develop/wanrun/internal/dog/core/handler"
	"github.com/wanrun-develop/wanrun/internal/router"
)

func Main() {
	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.CloseDB(dbConn)

	dogRepository := repository.NewDogRepository(dbConn)
	dogHandler := handler.NewDogHandler(dogRepository)
	dogController := controller.NewDogController(dogHandler)
	e := router.NewRouter(dogController)

	e.Logger.Fatal(e.Start(":8080"))
}
