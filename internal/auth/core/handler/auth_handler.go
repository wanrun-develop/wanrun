package handler

import (
	"github.com/wanrun-develop/wanrun/internal/auth/adapters/repository"
	model "github.com/wanrun-develop/wanrun/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type IAuthHandler interface {
	SignUp(authUser *model.AuthDogOwner) (model.ResAuthDogOwner, error)
	// LogIn() error
	// LogOut() error
}

type authHandler struct {
	ar repository.IAuthRepository
}

func NewAuthHandler(ar repository.IAuthRepository) IAuthHandler {
	return &authHandler{ar}
}

// SignUp
func (ah *authHandler) SignUp(authUser *model.AuthDogOwner) (model.ResAuthDogOwner, error) {
	// パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(authUser.Password), bcrypt.DefaultCost) // 一旦costをデフォルト値
	if err != nil {
		return model.ResAuthDogOwner{}, err
	}

	authUser.Password = string(hash)

	// ドッグのオーナー作成
	result, err := ah.ar.CreateDogOwner(authUser)

	resDogOwnerDetail := model.ResAuthDogOwner{
		AuthDogOwnerID: result.AuthDogOwnerID,
		Name:           result.DogOwner.Name,
		Email:          result.DogOwner.Email,
		Sex:            result.DogOwner.Sex,
	}

	return resDogOwnerDetail, err
}

func (ah *authHandler) LogIn() error  { return nil }
func (ah *authHandler) LogOut() error { return nil }
