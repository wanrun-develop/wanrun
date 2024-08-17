package handler

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/dogrun/adapters/googleplace"
)

type IDogrunHandler interface {
	GetDogrunDetail(c echo.Context, palceId string) (*googleplace.BaseResource, error)
}

type dogrunHandler struct {
	rest googleplace.IRest
}

func NewDogrunHandler(rest googleplace.IRest) IDogrunHandler {
	return &dogrunHandler{rest}
}

func (h *dogrunHandler) GetDogrunDetail(c echo.Context, palceId string) (*googleplace.BaseResource, error) {
	//base情報のFieldを使用
	var baseFiled googleplace.IFieldMask = googleplace.BaseField{}
	//place情報の取得
	res, err := h.rest.GETPlaceInfo(c, palceId, baseFiled)
	if err != nil {
		return nil, err
	}
	// JSONデータを構造体にデコード
	var apiResponse googleplace.BaseResource
	err = json.Unmarshal(res, &apiResponse)
	if err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
