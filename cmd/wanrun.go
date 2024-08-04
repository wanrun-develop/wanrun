package wanruncmd

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/db"
	"github.com/wanrun-develop/wanrun/internal/router"
)

func Main() {
	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.CloseDB(dbConn)
	e := echo.New()
	router.NewRouter(e, dbConn)

	e.Logger.Fatal(e.Start(":8080"))
}
