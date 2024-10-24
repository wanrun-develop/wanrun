package handler

import (
	_ "context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/configs"
	"github.com/wanrun-develop/wanrun/internal/auth/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/auth/core/dto"
	model "github.com/wanrun-develop/wanrun/internal/models"
	wrErrors "github.com/wanrun-develop/wanrun/pkg/errors"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"github.com/wanrun-develop/wanrun/pkg/success"
	wrUtil "github.com/wanrun-develop/wanrun/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

type IAuthHandler interface {
	SignUp(c echo.Context, reqADOD dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error)
	JwtProcessing(c echo.Context, rdo dto.ResDogOwnerDto) error
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
func (ah *authHandler) SignUp(c echo.Context, rado dto.ReqAuthDogOwnerDto) (dto.ResDogOwnerDto, error) {
	logger := log.GetLogger(c).Sugar()

	// パスワードのハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(rado.Password), bcrypt.DefaultCost) // 一旦costをデフォルト値

	if err != nil {
		wrErr := wrErrors.NewWRError(
			err,
			"パスワードに不正な文字列が入っております。",
			wrErrors.NewDogownerClientErrorEType(),
		)
		logger.Error(wrErr)
		return dto.ResDogOwnerDto{}, wrErr
	}

	// EmailとPhoneNumberのバリデーション
	if wrErr := validateEmailOrPhoneNumber(rado.Email, rado.PhoneNumber); wrErr != nil {
		logger.Error(wrErr)
		return dto.ResDogOwnerDto{}, wrErr
	}

	// requestからDogOwnerの構造体に詰め替え
	dogOwnerCredential := model.DogOwnerCredential{
		Email:       wrUtil.NewSqlNullString(rado.Email),
		PhoneNumber: wrUtil.NewSqlNullString(rado.PhoneNumber),
		Password:    wrUtil.NewSqlNullString(string(hash)),
		GrantType:   wrUtil.NewSqlNullString(model.PASSWORD_GRANT_TYPE), // Password認証
		AuthDogOwner: model.AuthDogOwner{
			DogOwner: model.DogOwner{
				Name: wrUtil.NewSqlNullString(rado.DogOwnerName),
			},
		},
	}

	logger.Debugf("dogOwnerCredential %v, Type: %T", dogOwnerCredential, dogOwnerCredential)

	// ドッグのオーナー作成
	result, wrErr := ah.ar.CreateDogOwner(c, &dogOwnerCredential)

	if wrErr != nil {
		return dto.ResDogOwnerDto{}, wrErr
	}

	// sessionIDの生成
	sessionID, wrErr := createSessionID(c, 15)

	if wrErr != nil {
		return dto.ResDogOwnerDto{}, wrErr
	}

	// 作成したDogOwnerの情報をdto詰め替え
	resDogOwnerDetail := dto.ResDogOwnerDto{
		DogOwnerID: uint64(result.AuthDogOwner.DogOwnerID.Int64),
		SessionID:  sessionID,
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

// validateEmailOrPhoneNumber: EmailかPhoneNumberの識別バリデーション。パスワード認証は、EmailかPhoneNumberで登録するため
//
// args:
//   - string: email Email
//   - string: phoneNumber 電話番号
//
// return:
//   - error: err
func validateEmailOrPhoneNumber(email string, phoneNumber string) error {
	// 両方が空の場合はエラー
	if email == "" && phoneNumber == "" {
		wrErr := wrErrors.NewWRError(
			nil,
			"Emailと電話番号のどちらも空です",
			wrErrors.NewDogownerClientErrorEType(),
		)
		return wrErr
	}

	// 両方に値が入っている場合もエラー
	if email != "" && phoneNumber != "" {
		wrErr := wrErrors.NewWRError(
			nil,
			"Emailと電話番号のどちらも値が入っています",
			wrErrors.NewDogownerClientErrorEType(),
		)
		return wrErr
	}

	// どちらか片方だけが入力されている場合は正常
	return nil
}

/*
jwt処理
*/
// JwtProcessing: jwtの生成等を行う
//
// args:
//   - echo.Context: c   Echoのコンテキスト。リクエストやレスポンスにアクセスするために使用されます。
//   - dto.ResDogOwnerDto: rdo フロントに返す飼い主情報
//
// return:
//  - error: error情報

func (ah *authHandler) JwtProcessing(c echo.Context, rdo dto.ResDogOwnerDto) error {
	// 秘密鍵取得
	secretKey := configs.FetchCondigStr("jwt.os.secret.key")
	jwtExpTime := configs.FetchCondigInt("jwt.exp.time")

	// jwt token生成
	signedToken, wrErr := createToken(c, secretKey, rdo, jwtExpTime)

	if wrErr != nil {
		return wrErr
	}

	return c.JSON(http.StatusCreated, success.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "飼い主の登録完了しました。",
		Token:   signedToken,
	})
}

// createToken: 指定された秘密鍵を使用して認証用のJWTトークンを生成します。
//
// args:
//   - echo.Context: c   Echoのコンテキスト。リクエストやレスポンスにアクセスするために使用されます
//   - string: secretKey   トークンの署名に使用する秘密鍵を表す文字列
//   - dto.ResDogOwnerDto: rdo 飼い主用のレスポンス情報
//   - int: expTime トークンの有効期限を秒単位で指定
//
// return:
//   - string: 生成されたJWTトークンを表す文字列。
//   - error: トークンの生成中に問題が発生した場合にはエラーを返します。
func createToken(c echo.Context, secretKey string, rdo dto.ResDogOwnerDto, expTime int) (string, error) {
	logger := log.GetLogger(c).Sugar()
	// JWTのペイロード
	claims := &dto.AccountClaims{
		ID:        strconv.FormatUint(uint64(rdo.DogOwnerID), 10), // stringにコンバート
		SessionID: rdo.SessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expTime))), // 有効時間
		},
	}

	// token生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// tokenに署名
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		wrErr := wrErrors.NewWRError(
			err,
			"パスワードに不正な文字列が入っております。",
			wrErrors.NewDogownerClientErrorEType(),
		)
		logger.Error(wrErr)
		return "", err
	}

	return signedToken, nil
}

// createSessionID: sessionIDの生成。引数の数だけランダムの文字列を生成する
//
// args:
//   - int: length 生成したい数
//
// return:
//   - string:　ランダム文字列
//   - error:　error情報
func createSessionID(c echo.Context, length int) (string, error) {
	logger := log.GetLogger(c).Sugar()

	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		wrErr := wrErrors.NewWRError(
			err,
			"SessionID生成に失敗しました。",
			wrErrors.NewDogownerServerErrorEType(),
		)
		logger.Error(wrErr)
		return "", wrErr
	}
	return base64.RawURLEncoding.EncodeToString(b)[:length], nil
}
