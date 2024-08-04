package wanruncmd

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/configs"
	"github.com/wanrun-develop/wanrun/internal/db"
	"github.com/wanrun-develop/wanrun/internal/dog/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/dog/controller"
	"github.com/wanrun-develop/wanrun/internal/dog/core/handler"
	"github.com/wanrun-develop/wanrun/internal/router"
)

func init() {
	if err := configs.LoadConfig(); err != nil {
		log.Fatal("設定ファイルのLoadに失敗しました。")
	}
}

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
