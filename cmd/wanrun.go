package wanruncmd

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/configs"
	"github.com/wanrun-develop/wanrun/internal/db"
)

func init() {
	if err := configs.LoadConfig(); err != nil {
		log.Fatal("設定ファイルのLoadに失敗しました。")
	}
}

func Main() {
	e := echo.New()

	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatalln(err)
	}

	defer db.CloseDB(dbConn)
	time := time.Now()
	message := fmt.Sprintf(
		"%v\nNowTime: %v",
		"Hello, World!!!!!",
		time)
	log.Printf("NowTime: %v\n", time)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
