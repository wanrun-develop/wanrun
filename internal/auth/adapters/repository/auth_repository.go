package repository

import (
	model "github.com/wanrun-develop/wanrun/internal/models"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateDogOwner(authDogOwner *model.AuthDogOwner) (*model.AuthDogOwner, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{db}
}

func (ar *authRepository) CreateDogOwner(authDogOwner *model.AuthDogOwner) (*model.AuthDogOwner, error) {
	err := ar.db.Transaction(func(tx *gorm.DB) error {
		// DogOwnerのレコード作成
		if err := tx.Create(&authDogOwner.DogOwner).Error; err != nil {
			return err
		}
		// AuthDogOwnerのレコード作成
		if err := tx.Create(&authDogOwner).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	result := model.AuthDogOwner{}

	// レスポンス用にDogOwner情報の取得
	err = ar.db.Preload("DogOwner").First(&result, authDogOwner.AuthDogOwnerID).Error

	if err != nil {
		return nil, err
	}
	return &result, nil
}
