package handler

import (
	"github.com/wanrun-develop/wanrun/internal/dog/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/dog/core/model"
)

type IDogHandler interface {
	GetAllDogs() ([]model.DogRes, error)
}

type dogHandler struct {
	dr repository.IDogRepository
}

func NewDogHandler(dr repository.IDogRepository) IDogHandler {
	return &dogHandler{dr}
}

func (dh *dogHandler) GetAllDogs() ([]model.DogRes, error) {
	dogs, err := dh.dr.GetAllDogs()

	if err != nil {
		return []model.DogRes{}, err
	}

	resDogs := []model.DogRes{}

	for _, d := range dogs {
		dr := model.DogRes{
			DogID:      d.DogID,
			DogOwnerID: d.DogOwnerID,
			Name:       d.Name,
			DogTypeID:  d.DogTypeID,
			Weight:     d.Weight,
			Sex:        d.Sex,
			Image:      d.Image,
			CreateAt:   d.CreateAt,
			UpdateAt:   d.UpdateAt,
		}
		resDogs = append(resDogs, dr)
	}
	return resDogs, nil
}
