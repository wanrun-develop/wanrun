package repository

import "gorm.io/gorm"

type IAuthRepository interface {
	CreateDogOwner() error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{db}
}

func (ar *authRepository) CreateDogOwner() error { return nil }
