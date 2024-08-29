package handler

import (
	"github.com/wanrun-develop/wanrun/internal/auth/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/auth/core/dto"
	model "github.com/wanrun-develop/wanrun/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type IAuthHandler interface {
	SignUp(authUser dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error)
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
func (ah *authHandler) SignUp(reqADOD dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error) {
	// パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(reqADOD.Password), bcrypt.DefaultCost) // 一旦costをデフォルト値
	if err != nil {
		return dto.ResDogOwnerDto{}, err
	}

	// requestからauthDogOwnerの構造体に詰め替え
	authDogOwner := model.AuthDogOwner{
		Password: string(hash),
		DogOwner: model.DogOwner{
			Name:  reqADOD.DogOwnerName,
			Email: reqADOD.Email,
		},
	}

	// ドッグのオーナー作成
	result, err := ah.ar.CreateDogOwner(&authDogOwner)

	// 作成したDogOwnerの情報を詰め替え
	resDogOwnerDetail := dto.ResDogOwnerDto{
		DogOwnerID: result.AuthDogOwnerID,
	}
	return resDogOwnerDetail, err
}

func (ah *authHandler) LogIn() error  { return nil }
func (ah *authHandler) LogOut() error { return nil }
