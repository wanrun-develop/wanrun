package handler

import (
	_ "context"
	_ "errors"
	_ "time"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/auth/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/auth/core/dto"
	model "github.com/wanrun-develop/wanrun/internal/models"
	"github.com/wanrun-develop/wanrun/internal/models/types"
	_ "github.com/wanrun-develop/wanrun/internal/models/types"

	// _ wrErrors "github.com/wanrun-develop/wanrun/pkg/errors"
	"github.com/wanrun-develop/wanrun/pkg/log"
	wrUtil "github.com/wanrun-develop/wanrun/pkg/util"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

type IAuthHandler interface {
	SignUp(c echo.Context, reqADOD dto.ReqAuthDogOwnerDto, grantType types.GrantType) (dto.ResDogOwnerDto, error)
	// LogIn(c echo.Context, reqADOD dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error)
	// LogOut() error
	// GoogleOAuth(c echo.Context, authorizationCode string, grantType types.GrantType) (dto.ResDogOwnerDto, error)
}

type authHandler struct {
	ar repository.IAuthRepository
	// ag google.IOAuthGoogle
}

//	func NewAuthHandler(ar repository.IAuthRepository, g google.IOAuthGoogle) IAuthHandler {
//		return &authHandler{ar, g}
//	}
func NewAuthHandler(ar repository.IAuthRepository) IAuthHandler {
	return &authHandler{ar}
}

// SignUp
func (ah *authHandler) SignUp(c echo.Context, rado dto.ReqAuthDogOwnerDto, gt types.GrantType) (dto.ResDogOwnerDto, error) {
	logger := log.GetLogger(c).Sugar()

	// パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(rado.Password), bcrypt.DefaultCost) // 一旦costをデフォルト値

	if err != nil {
		logger.Error(err)
		return dto.ResDogOwnerDto{}, err
	}

	// EmailとPhoneNumberのバリデーション
	if wrErr := dto.ValidateEmailOrPhoneNumber(rado.Email, rado.PhoneNumber); wrErr != nil {
		logger.Error(wrErr)
		return dto.ResDogOwnerDto{}, wrErr
	}

	// requestからauthDogOwnerの構造体に詰め替え
	dogOwnerCredential := model.DogOwnerCredential{
		Email:       wrUtil.NewSqlNullString(rado.Email),
		PhoneNumber: wrUtil.NewSqlNullString(rado.PhoneNumber),
		Password:    wrUtil.NewSqlNullString(string(hash)),
		AuthDogOwner: model.AuthDogOwner{
			GrantType: gt,
			DogOwner: model.DogOwner{
				Name: wrUtil.NewSqlNullString(rado.DogOwnerName),
			},
		},
	}

	logger.Debugf("dogOwnerCredential %v, Type: %T", dogOwnerCredential, dogOwnerCredential)

	// ドッグのオーナー作成
	result, err := ah.ar.CreateDogOwner(c, &dogOwnerCredential)

	if err != nil {
		return dto.ResDogOwnerDto{}, err
	}

	// 作成したDogOwnerの情報を詰め替え
	resDogOwnerDetail := dto.ResDogOwnerDto{
		DogOwnerID: uint64(result.AuthDogOwner.DogOwnerID.Int64),
	}

	logger.Infof("resDogOwnerDetail: %v", resDogOwnerDetail)

	return resDogOwnerDetail, nil
}

// Login
// func (ah *authHandler) LogIn(c echo.Context, reqADOD dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error) {
// 	logger := log.GetLogger(c).Sugar()
// 	authDogOwner := model.AuthDogOwner{
// 		DogOwner: model.DogOwner{
// 			Email: reqADOD.Email,
// 		},
// 	}

// 	logger.Infof("authDogOwner Info: %v", authDogOwner)

// 	// Emailから対象のDogOwner情報の取得
// 	result, err := ah.ar.GetDogOwnerByEmail(c, authDogOwner)

// 	if err != nil {
// 		logger.Error(err)
// 		return dto.ResDogOwnerDto{}, err
// 	}

// 	// パスワードの確認
// 	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(reqADOD.Password))

// 	if err != nil {
// 		logger.Error(err)
// 		return dto.ResDogOwnerDto{}, err
// 	}

// 	resDogOwnerDetail := dto.ResDogOwnerDto{
// 		DogOwnerID: result.DogOwnerID,
// 	}

// 	logger.Infof("resDogOwnerDetail: %v", resDogOwnerDetail)

// 	return resDogOwnerDetail, nil
// }

// Logout
func (ah *authHandler) LogOut() error { return nil }

/*
Google OAuth認証
*/
// func (ah *authHandler) GoogleOAuth(c echo.Context, authorizationCode string, grantType types.GrantType) (dto.ResDogOwnerDto, error) {
// 	logger := log.GetLogger(c).Sugar()

// 	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second) // 5秒で設定
// 	defer cancel()

// 	// 各token情報の取得
// 	token, wrErr := ah.ag.GetAccessToken(c, authorizationCode, ctx)

// 	if wrErr != nil {
// 		return dto.ResDogOwnerDto{}, wrErr
// 	}

// 	// トークン元にGoogleユーザー情報の取得
// 	googleUserInfo, wrErr := ah.ag.GetGoogleUserInfo(c, token, ctx)

// 	if wrErr != nil {
// 		return dto.ResDogOwnerDto{}, wrErr
// 	}

// 	// Googleユーザー情報の確認処理
// 	if googleUserInfo == nil {
// 		wrErr := wrErrors.NewWRError(
// 			errors.New(""),
// 			"no google user information",
// 			wrErrors.NewAuthServerErrorEType(),
// 		)
// 		logger.Errorf("No google user information error: %v", wrErr)
// 		return dto.ResDogOwnerDto{}, wrErr
// 	}

// 	// ドッグオーナーのcredentialの設定と型変換
// 	dogOwnerCredential := model.DogOwnerCredential{
// 		ProviderUserID: wrUtil.NewSqlNullString(googleUserInfo.UserId),
// 		Email:          wrUtil.NewSqlNullString(googleUserInfo.Email),
// 		AuthDogOwner: model.AuthDogOwner{
// 			AccessToken:           wrUtil.NewSqlNullString(token.AccessToken),
// 			RefreshToken:          wrUtil.NewSqlNullString(token.RefreshToken),
// 			AccessTokenExpiration: wrUtil.NewCustomTime(token.Expiry),
// 			GrantType:             grantType,
// 			DogOwner: model.DogOwner{
// 				Name: wrUtil.NewSqlNullString(googleUserInfo.Email),
// 			},
// 		},
// 	}

// 	// ドッグオーナーの作成
// 	dogOC, wrErr := ah.ar.CreateOAuthDogOwner(c, &dogOwnerCredential)

// 	if wrErr != nil {
// 		return dto.ResDogOwnerDto{}, wrErr
// 	}

// 	resDogOwner := dto.ResDogOwnerDto{
// 		DogOwnerID: uint(dogOC.AuthDogOwner.DogOwner.DogOwnerID.Int64),
// 	}

// 	return resDogOwner, nil
// }
