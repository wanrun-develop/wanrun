package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/common"
	"github.com/wanrun-develop/wanrun/internal/dogrun/core/dto"
	"github.com/wanrun-develop/wanrun/internal/dogrun/core/handler"
	"github.com/wanrun-develop/wanrun/pkg/errors"
	"github.com/wanrun-develop/wanrun/pkg/log"
)

type IDogrunController interface {
	GetDogrunDetail(echo.Context) error
	GetDogrunDetailByID(echo.Context) error
	GetDogrunTagMst(echo.Context) error
	SearchAroundDogruns(echo.Context) error
	GetDogrunPhoto(echo.Context) error
	GetBookmarkedDogruns(echo.Context) error
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

	dogrun, err := dc.h.GetDogrunDetail(c, placeId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dogrun)
}

// GetDogrunDetailByID: DogrunIDによるドッグラン詳細情報の取得
//
// args:
//   - echo.Context:	コンテキスト
//
// return:
//   - error:	エラー
func (dc *dogrunController) GetDogrunDetailByID(c echo.Context) error {
	logger := log.GetLogger(c).Sugar()

	dogrunIDStr := c.Param("dogrunId")
	logger.Info("リクエストdogrun id :", dogrunIDStr)

	// ドッグランIDのバリデーション
	if dogrunIDStr == "" {
		err := errors.NewWRError(nil, "dogrun idが指定されていません", errors.NewDogrunClientErrorEType())
		logger.Error(err)
		return err
	}

	// 数値形式チェック
	dogrunID, err := strconv.ParseInt(dogrunIDStr, 10, 64)
	if err != nil {
		wrErr := errors.NewWRError(err, "dogrun idの形式が不正です", errors.NewDogrunClientErrorEType())
		logger.Error(wrErr)
		return wrErr
	}

	dogrun, err := dc.h.GetDogrunDetailByID(c, dogrunID)
	if err != nil {
		// すでにWRErrorである場合は変換せずにログを出力して返す
		logger.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, dogrun)
}

// GetDogrunTagMst: DogrunTagMstのマスターデータの取得
//
// args:
//   - echo.Context:	コンテキスト
//
// return:
//   - error:	エラー
func (dc *dogrunController) GetDogrunTagMst(c echo.Context) error {
	logger := log.GetLogger(c).Sugar()
	logger.Info("dogrun tagMst情報の取得開始")

	mstRes, err := dc.h.GetDogrunTagMst(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mstRes)
}

// ドッグランの周辺検索
func (dc *dogrunController) SearchAroundDogruns(c echo.Context) error {
	logger := log.GetLogger(c).Sugar()
	//リクエストボディをバインド
	var condition dto.SearchAroundRectangleCondition
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

	if len(condition.IncludeDogrunTags) == 0 {
		//ドッグランタグ検索条件がない場合は、google検索を基準とする
		if resDogruns, err := dc.h.SearchAroundDogruns(c, condition); err != nil {
			return err
		} else {
			return c.JSON(http.StatusOK, resDogruns)

		}
	} else {
		//ドッグランタグ検索条件がある場合は、DB検索を基準とする
		if resDogruns, err := dc.h.SearchAroundAndTagDogruns(c, condition); err != nil {
			return err
		} else {
			return c.JSON(http.StatusOK, resDogruns)

		}
	}

}

// ドッグランの画像nameよりsrcUrlの取得
func (dc *dogrunController) GetDogrunPhoto(c echo.Context) error {
	logger := log.GetLogger(c).Sugar()

	name := c.QueryParam("name")
	widthPx := c.QueryParam("widthPx")
	err := validateMaxPX(widthPx)
	if err != nil {
		logger.Error(err)
		return err
	}
	heightPx := c.QueryParam("heightPx")
	err = validateMaxPX(heightPx)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("photo name :", name)
	logger.Info("photo widthPx :", widthPx)
	logger.Info("photo heightPx :", heightPx)

	srcUri, err := dc.h.GetDogrunPhotoSrc(c, name, widthPx, heightPx)
	if err != nil {
		logger.Error(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"src": srcUri,
	})
}

/*
リクエストのクエリパラメータのpxのバリデーション
*/
func validateMaxPX(px string) error {
	// stringをintに変換
	convertedPX, err := strconv.Atoi(px)
	if err != nil {
		return errors.NewWRError(nil, "リクエストの画像サイズの指定が不正です。", errors.NewDogrunClientErrorEType())
	}

	// 1から4800の範囲であるかをチェック
	if convertedPX < 1 || convertedPX > 4800 {
		return errors.NewWRError(nil, "リクエストの画像サイズの指定が不正です。1以上4800以下である必要があります。", errors.NewDogrunClientErrorEType())
	}
	return nil
}

// GetBookmarkedDogruns: ブックマークしたドッグランの一覧取得
//
// args:
//   - echo.Context:	コンテキスト
//
// return:
// error:	エラー
func (dc *dogrunController) GetBookmarkedDogruns(c echo.Context) error {
	logger := log.GetLogger(c).Sugar()

	pagination := common.PaginationReq{}
	if err := c.Bind(&pagination); err != nil {
		err = errors.NewWRError(err, "paginationが不正です", errors.NewDogrunClientErrorEType())
		logger.Error(err)
		return err
	}

	// バリデータのインスタンス作成
	validate := validator.New()
	//リクエストボディのバリデーション
	if err := validate.Struct(pagination); err != nil {
		err = errors.NewWRError(err, "paginationのバリデーションに違反しています", errors.NewDogrunClientErrorEType())
		logger.Error(err)
		return err
	}

	dogruns, err := dc.h.GetBookmarkedDogruns(c, pagination)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, dogruns)
}
