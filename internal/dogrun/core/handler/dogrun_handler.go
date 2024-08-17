package handler

import (
	"encoding/json"

	"github.com/wanrun-develop/wanrun/internal/dogrun/adapters/googleplace"
)

type IDogrunHandler interface {
	GetDogrunDetail(palceId string) (*googleplace.BaseResource, error)
}

type dogrunHandler struct {
	rest googleplace.IRest
}

func NewDogrunHandler(rest googleplace.IRest) IDogrunHandler {
	return &dogrunHandler{rest}
}

func (h *dogrunHandler) GetDogrunDetail(palceId string) (*googleplace.BaseResource, error) {
	var baseFiled googleplace.IFieldMask = googleplace.BaseField{}
	res, err := h.rest.GETPlaceInfo(palceId, baseFiled)
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
