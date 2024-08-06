package repository

import (
	"fmt"

	"github.com/wanrun-develop/wanrun/internal/dog/core/model"
	"gorm.io/gorm"
)

type IDogRepository interface {
	GetAllDogs() ([]model.Dog, error)
	GetDogByID(dogID uint) (model.Dog, error)
	CreateDog() (model.Dog, error)
	DeleteDog(dogID uint) error
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
	if err := dr.db.Where("dog_id=?", dogID).First(&dog).Error; err != nil {
		return model.Dog{}, err
	}
	return dog, nil
}

func (dr *dogRepository) CreateDog() (model.Dog, error) {
	dog := model.Dog{}
	if err := dr.db.Create(&dog).Error; err != nil {
		return model.Dog{}, err
	}
	return dog, nil
}

func (dr *dogRepository) DeleteDog(dogID uint) error {
	result := dr.db.Where("dog_id=?", dogID).Delete(&model.Dog{})

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

// func (dr *dogRepository) UpdateDog(dogID uint) error {
// 	dog := model.Dog{}
// 	result := dr.db.Model(&dog).Clauses(clause.Returning{}).Where("dog_id=?", dogID).Update()
// }
