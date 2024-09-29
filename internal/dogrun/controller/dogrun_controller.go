package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/dogrun/core/dto"
	"github.com/wanrun-develop/wanrun/internal/dogrun/core/handler"
	"github.com/wanrun-develop/wanrun/pkg/errors"
	"github.com/wanrun-develop/wanrun/pkg/log"
)

type IDogrunController interface {
	GetDogrunDetail(echo.Context) error
	GetDogrun(echo.Context) error
	SearchAroundDogruns(echo.Context) error
}

type dogrunController struct {
	h handler.IDogrunHandler
}

func NewDogrunController(h handler.IDogrunHandler) IDogrunController {
	return &dogrunController{h}
}

// ドッグラン詳細情報の取得
func (dc *dogrunController) GetDogrunDetail(c echo.Context) error {
	logger := log.GetLogger(c).Sugar()

	placeId := c.Param("placeId")
	logger.Info("リクエストplace id :", placeId)
	log.GetLogger(c).Info("リクエストplace id")

	dogrun, err := dc.h.GetDogrunDetail(c, placeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve dog run information",
		})
	}

	return c.JSON(http.StatusOK, dogrun)
}

func (dc *dogrunController) GetDogrun(c echo.Context) error {
	id := c.Param("id")
	dc.h.GetDogrunByID(id)
	return nil
}

func (dc *dogrunController) SearchAroundDogruns(c echo.Context) error {
	logger := log.GetLogger(c).Sugar()
	//リクエストボディをバインド
	var condition dto.SearchAroudRectangleCondition
	if err := c.Bind(&condition); err != nil {
		err = errors.NewWRError(err, "検索条件が不正です", errors.NewDogrunClientErrorEType())
		logger.Error(err)
		return err
	}
	// バリデータのインスタンス作成
	validate := validator.New()
	// カスタムバリデーションルールの登録
	_ = validate.RegisterValidation("latitude", dto.VLatitude)
	_ = validate.RegisterValidation("longitude", dto.VLongitude)

	//リクエストボディのバリデーション
	if err := validate.Struct(condition); err != nil {
		err = errors.NewWRError(err, "検索条件のバリデーションに違反しています", errors.NewDogrunClientErrorEType())
		logger.Error(err)
		return err
	}

	dogruns, err := dc.h.SearchAroundDogruns(c, condition)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve dog run information",
		})
	}
	return c.JSON(http.StatusOK, dogruns)
}
