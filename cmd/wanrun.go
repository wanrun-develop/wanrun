package wanruncmd

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/wanrun-develop/wanrun/internal/db"
)

func Main() {
	e := echo.New()

	// db := db.NewDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!!!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
