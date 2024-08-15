package wanruncmd

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wanrun-develop/wanrun/configs"
	authRepository "github.com/wanrun-develop/wanrun/internal/auth/adapters/repository"
	authController "github.com/wanrun-develop/wanrun/internal/auth/controller"
	authHandler "github.com/wanrun-develop/wanrun/internal/auth/core/handler"
	"github.com/wanrun-develop/wanrun/internal/db"
	dogRepository "github.com/wanrun-develop/wanrun/internal/dog/adapters/repository"
	dogController "github.com/wanrun-develop/wanrun/internal/dog/controller"
	dogHandler "github.com/wanrun-develop/wanrun/internal/dog/core/handler"
	logger "github.com/wanrun-develop/wanrun/pkg/log"
	"gorm.io/gorm"
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

	// グローバルロガーの初期化
	zap := logger.NewWanRunLogger()
	logger.SetLogger(zap) // グローバルロガーを設定
	// アプリケーション終了時にロガーを同期
	defer zap.Sync()

	// ミドルウェアを登録
	e.Use(middleware.RequestID())
	e.Use(logger.RequestLoggerMiddleware(zap))

	// Router設定
	newRouter(e, dbConn)
	e.GET("/test", logger.Test)

	e.Logger.Fatal(e.Start(":8080"))
}

func newRouter(e *echo.Echo, dbConn *gorm.DB) {
	dogController := newDog(dbConn)

	// e.POST("/sign-up")
	dog := e.Group("dog")
	dog.GET("/all", dogController.GetAllDogs)
	dog.GET("/:dogID", dogController.GetDogByID)
	dog.POST("/create", dogController.CreateDog)
	dog.DELETE("/delete", dogController.DeleteDog)
	// dog.PUT("/:dogID", dogController.UpdateDog)

	authController := newAuth(dbConn)
	auth := e.Group("auth")
	auth.POST("/signUp", authController.SignUp)
}

// dogの初期化
func newDog(dbConn *gorm.DB) dogController.IDogController {
	dogRepository := dogRepository.NewDogRepository(dbConn)
	dogHandler := dogHandler.NewDogHandler(dogRepository)
	dogController := dogController.NewDogController(dogHandler)

	return dogController
}

func newAuth(dbConn *gorm.DB) authController.IAuthController {
	authRepository := authRepository.NewAuthRepository(dbConn)
	authHandler := authHandler.NewAuthHandler(authRepository)
	authController := authController.NewAuthController(authHandler)

	return authController
}
