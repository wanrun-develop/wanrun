package handler

import (
	"github.com/wanrun-develop/wanrun/internal/dog/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/dog/core/model"
)

type IDogHandler interface {
	GetAllDogs() ([]model.DogRes, error)
	GetDogByID(dogID uint) (model.DogRes, error)
}

type dogHandler struct {
	dr repository.IDogRepository
}

func NewDogHandler(dr repository.IDogRepository) IDogHandler {
	return &dogHandler{dr}
}

func (dh *dogHandler) GetAllDogs() ([]model.DogRes, error) {
	dogs := []model.Dog{}

	if err := dh.dr.GetAllDogs(&dogs); err != nil {
		return nil, err
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

func (dh *dogHandler) GetDogByID(dogID uint) (model.DogRes, error) {
	dog := model.Dog{}

	if err := dh.dr.GetDogByID(&dog, dogID); err != nil {
		return model.DogRes{}, err
	}

	resDog := model.DogRes{
		DogID:      dog.DogID,
		DogOwnerID: dog.DogOwnerID,
		Name:       dog.Name,
		DogTypeID:  dog.DogTypeID,
		Weight:     dog.Weight,
		Sex:        dog.Sex,
		Image:      dog.Image,
		CreateAt:   dog.CreateAt,
		UpdateAt:   dog.UpdateAt,
	}
	return resDog, nil
}
