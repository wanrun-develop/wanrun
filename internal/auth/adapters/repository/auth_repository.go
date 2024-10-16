package repository

import (
	"errors"

	"github.com/labstack/echo/v4"
	model "github.com/wanrun-develop/wanrun/internal/models"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateDogOwner(c echo.Context, authDogOwner *model.AuthDogOwner) (*model.AuthDogOwner, error)
	GetDogOwnerByEmail(c echo.Context, authDogOwner model.AuthDogOwner) (*model.AuthDogOwner, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{db}
}

func (ar *authRepository) CreateDogOwner(c echo.Context, authDogOwner *model.AuthDogOwner) (*model.AuthDogOwner, error) {
	logger := log.GetLogger(c).Sugar()

	// Email重複のvalidate
	var emailExistingCount int64
	err := ar.db.Model(&model.DogOwner{}).Where("email = ?", authDogOwner.DogOwner.Email).Count(&emailExistingCount).Error

	if err != nil {
		logger.Error("Failed to check existing email: ", err)
		return nil, err
	}
	if emailExistingCount > 0 {
		logger.Info("Email already exists")
		return nil, errors.New("email already exists")
	}

	// トランザクションの開始
	err = ar.db.Transaction(func(tx *gorm.DB) error {
		// DogOwnerのレコード作成
		if err := tx.Create(&authDogOwner.DogOwner).Error; err != nil {
			logger.Error("Failed to DogOwner: ", err)
			return err
		}

		// DogOwnerが作成された後、そのIDをauthDogOwnerに設定
		authDogOwner.DogOwnerID = authDogOwner.DogOwner.DogOwnerID

		// AuthDogOwnerのレコード作成
		if err := tx.Create(&authDogOwner).Error; err != nil {
			logger.Error("Failed to AuthDogOwner: ", err)
			return err
		}
		return nil
	})

	if err != nil {
		logger.Error("Transaction failed: ", err)
		return nil, err
	}

	logger.Infof("Created DogOwner Detail: %v", authDogOwner)

	result := model.AuthDogOwner{}

	// レスポンス用にDogOwner情報の取得
	err = ar.db.Preload("DogOwner").First(&result, authDogOwner.AuthDogOwnerID).Error

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Infof("AuthDogOwner Info: %v", result)

	return &result, nil
}

/*
DogOwnerのメールアドレスからDogOwner情報の取得
*/
func (ar *authRepository) GetDogOwnerByEmail(c echo.Context, authDogOwner model.AuthDogOwner) (*model.AuthDogOwner, error) {
	logger := log.GetLogger(c).Sugar()

	var result model.AuthDogOwner = model.AuthDogOwner{}

	if err := ar.db.Preload("DogOwner").Joins("LEFT JOIN dog_owners ON dog_owners.dog_owner_id = auth_dog_owners.dog_owner_id").Where("dog_owners.email = ?", authDogOwner.DogOwner.Email).First(&result).Error; err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Infof("Query Result: %v\n", result)

	return &result, nil
}
