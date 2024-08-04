package repository

import (
	"github.com/wanrun-develop/wanrun/internal/dog/core/model"
	"gorm.io/gorm"
)

type IDogRepository interface {
	GetAllDogs() ([]model.Dog, error)
	GetDogByID(dogID uint) (model.Dog, error)
}

type dogRepository struct {
	db *gorm.DB
}

func NewDogRepository(db *gorm.DB) IDogRepository {
	return &dogRepository{db}
}

func (dr *dogRepository) GetAllDogs() ([]model.Dog, error) {
	dogs := []model.Dog{}
	if err := dr.db.Find(&dogs).Error; err != nil {
		return []model.Dog{}, err
	}
	return dogs, nil
}

func (dr *dogRepository) GetDogByID(dogID uint) (model.Dog, error) {
	dog := model.Dog{}
	if err := dr.db.Where("dog_id = ?", dogID).First(&dog).Error; err != nil {
		return model.Dog{}, err
	}
	return dog, nil
}
