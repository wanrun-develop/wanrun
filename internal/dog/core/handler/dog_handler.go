package handler

import (
	"github.com/wanrun-develop/wanrun/internal/dog/adapters/repository"
	model "github.com/wanrun-develop/wanrun/internal/models"
)

type IDogHandler interface {
	GetAllDogs() ([]model.DogRes, error)
	GetDogByID(dogID uint) (model.DogRes, error)
	CreateDog() (model.DogRes, error)
	DeleteDog(dogID uint) error
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

func (dh *dogHandler) GetDogByID(dogID uint) (model.DogRes, error) {

	dog, err := dh.dr.GetDogByID(dogID)

	if err != nil {
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

func (dh *dogHandler) CreateDog() (model.DogRes, error) {
	dog, err := dh.dr.CreateDog()

	if err != nil {
		return model.DogRes{}, err
	}

	dogRes := model.DogRes{
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
	return dogRes, err
}

func (dh *dogHandler) DeleteDog(dogID uint) error {
	if err := dh.dr.DeleteDog(dogID); err != nil {
		return err
	}
	return nil
}
