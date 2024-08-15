package repository

import (
	model "github.com/wanrun-develop/wanrun/internal/models"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateDogOwner(authDogOwner *model.AuthDogOwner) (*model.ResAuthDogOwner, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{db}
}

func (ar *authRepository) CreateDogOwner(authDogOwner *model.AuthDogOwner) (*model.ResAuthDogOwner, error) {
	err := ar.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(authDogOwner).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	result := model.ResAuthDogOwner{}

	err = ar.db.Preload("DogOwner").First(&result, authDogOwner.AuthDogOwnerID).Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}
