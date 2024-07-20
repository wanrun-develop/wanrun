package wanruncmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/db"
)

func Main() {
	e := echo.New()

	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("DB info: %v", dbConn)

	defer db.CloseDB(dbConn)

	message := fmt.Sprintf("%v", "Hello, World!!!!!!")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
