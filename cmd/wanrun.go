package wanruncmd

import (
	"log"

	"github.com/labstack/echo/v4/middleware"
	"github.com/wanrun-develop/wanrun/configs"
	"github.com/wanrun-develop/wanrun/internal/db"
	"github.com/wanrun-develop/wanrun/internal/router"
	logger "github.com/wanrun-develop/wanrun/pkg/log"
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
	e := echo.New()
	router.NewRouter(e, dbConn)

	// グローバルロガーの初期化
	zap := logger.NewWanRunLogger()
	logger.SetLogger(zap) // グローバルロガーを設定
	// アプリケーション終了時にロガーを同期
	defer zap.Sync()

	// ミドルウェアを登録
	e.Use(middleware.RequestID())
	e.Use(logger.RequestLoggerMiddleware(zap))

	e.GET("/test", logger.Test)

	e.Logger.Fatal(e.Start(":8080"))
}
