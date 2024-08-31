package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/auth/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/auth/core/dto"
	model "github.com/wanrun-develop/wanrun/internal/models"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

type IAuthHandler interface {
	SignUp(c echo.Context, reqADOD dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error)
	LogIn(c echo.Context, reqADOD dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error)
	// LogOut() error
}

type authHandler struct {
	ar repository.IAuthRepository
}

func NewAuthHandler(ar repository.IAuthRepository) IAuthHandler {
	return &authHandler{ar}
}

// SignUp
func (ah *authHandler) SignUp(c echo.Context, reqADOD dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error) {
	logger := log.GetLogger(c).Sugar()
	// パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(reqADOD.Password), bcrypt.DefaultCost) // 一旦costをデフォルト値
	if err != nil {
		logger.Error(err)
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
	result, err := ah.ar.CreateDogOwner(c, &authDogOwner)

	if err != nil {
		logger.Error(err)
		return dto.ResDogOwnerDto{}, err
	}

	// 作成したDogOwnerの情報を詰め替え
	resDogOwnerDetail := dto.ResDogOwnerDto{
		DogOwnerID: result.AuthDogOwnerID,
	}

	logger.Infof("resDogOwnerDetail: %v", resDogOwnerDetail)

	return resDogOwnerDetail, nil
}

// Login
func (ah *authHandler) LogIn(c echo.Context, reqADOD dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error) {
	logger := log.GetLogger(c).Sugar()
	authDogOwner := model.AuthDogOwner{
		DogOwner: model.DogOwner{
			Email: reqADOD.Email,
		},
	}

	logger.Infof("authDogOwner Info: %v\n", authDogOwner)

	// Emailから対象のDogOwner情報の取得
	result, err := ah.ar.GetDogOwnerByEmail(c, authDogOwner)

	if err != nil {
		logger.Error(err)
		return dto.ResDogOwnerDto{}, err
	}

	// パスワードの確認
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(reqADOD.Password))

	if err != nil {
		logger.Error(err)
		return dto.ResDogOwnerDto{}, err
	}

	resDogOwnerDetail := dto.ResDogOwnerDto{
		DogOwnerID: result.DogOwnerID,
	}

	logger.Infof("resDogOwnerDetail: %v", resDogOwnerDetail)

	return resDogOwnerDetail, nil
}

// Logout
func (ah *authHandler) LogOut() error { return nil }
