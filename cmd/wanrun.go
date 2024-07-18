package wanruncmd

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/db"
)

func Main() {
	e := echo.New()

	db := db.NewDB()
	fmt.Printf("DB info: %v", db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!!!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
