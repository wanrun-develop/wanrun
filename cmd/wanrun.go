package wanruncmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/db"
	"github.com/wanrun-develop/wanrun/internal/dogOwner/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/dogOwner/core"
	"gorm.io/gorm"
)

func Main() {
	e := echo.New()

	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("DB info: %v", dbConn)

	defer db.CloseDB(dbConn)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!!!")
	})

	rooting(e, dbConn)

	e.Logger.Fatal(e.Start(":8080"))
}

/*
	ルーティング設定
*/
func rooting(e *echo.Echo, dbConn *gorm.DB) {

	//repository作成
	dogOwnerRepository := repository.NewDogOwnerRepository(dbConn)
	//handler作成
	dogOwnerhandler := core.NewDogOwnerHandler(dogOwnerRepository)

	dogOwnerG := e.Group("/dogowner")
	dogOwnerG.POST("/create", dogOwnerhandler.CreateDogOwner)
}
