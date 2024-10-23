package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/configs"
	"github.com/wanrun-develop/wanrun/internal/auth/core/dto"
	"github.com/wanrun-develop/wanrun/internal/auth/core/handler"
	"github.com/wanrun-develop/wanrun/internal/models/types"
	_ "github.com/wanrun-develop/wanrun/internal/models/types"
	"github.com/wanrun-develop/wanrun/pkg/errors"
	_ "github.com/wanrun-develop/wanrun/pkg/errors"
	wrErrors "github.com/wanrun-develop/wanrun/pkg/errors"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"github.com/wanrun-develop/wanrun/pkg/success"
)

type IAuthController interface {
	SignUp(c echo.Context) error
	// LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	// GoogleOAuth(c echo.Context) error
}

type authController struct {
	ah handler.IAuthHandler
}

func NewAuthController(ah handler.IAuthHandler) IAuthController {
	return &authController{ah}
}

/*
GoogleのOAuth認証
*/
// func (ac *authController) GoogleOAuth(c echo.Context) error {
// 	logger := log.GetLogger(c).Sugar()

// 	// GrantTypeヘッダーを取得
// 	grantTypeHeader := c.Request().Header.Get(string(types.OAUTH_IDENTIFICATION_HEADER))

// 	// GrantTypeヘッダーのバリデーション
// 	if err := dto.ValidateGrantTypeHeader(grantTypeHeader, string(types.OAUTH_GRANT_TYPE_HEADER)); err != nil {
// 		err = wrErrors.NewWRError(err, "ヘッダー情報が異なります。", wrErrors.NewDogrunClientErrorEType())
// 		logger.Error(err)
// 		return err
// 	}

// 	// GrantTypeに型変換
// 	grantType := types.GrantType(grantTypeHeader)
// 	logger.Infof("grantTypeHeader: %v", grantType)

// 	// 認証コードの取得
// 	authorizationCode := c.QueryParam("code")

// 	// ユーザーが承認しなかったら、エラーのクエリパラメータにくるため
// 	oauthErrorCode := c.QueryParam("error")

// 	logger.Infof("authorizationCode: %v, oauthErrorCode: %v", authorizationCode, oauthErrorCode)

// 	// クエリパラメータのバリデーション
// 	if err := dto.ValidateOAuthResCode(authorizationCode, oauthErrorCode); err != nil {
// 		err = wrErrors.NewWRError(err, "承認をしてください。", wrErrors.NewDogrunClientErrorEType())
// 		logger.Error(err)
// 		return err
// 	}

// 	// OAuth認証
// 	resDogOwner, wrErr := ac.ah.GoogleOAuth(c, authorizationCode, grantType)

// 	if wrErr != nil {
// 		return wrErr
// 	}

// 	// jwt処理
// 	return jwtProcessing(c, resDogOwner)
// }

/*
パスワード認証
*/
func (ac *authController) SignUp(c echo.Context) error {
	logger := log.GetLogger(c).Sugar()

	reqADOD := dto.ReqAuthDogOwnerDto{}

	if err := c.Bind(&reqADOD); err != nil {
		wrErr := errors.NewWRError(err, "入力項目に不正があります。", errors.NewDogrunClientErrorEType())
		logger.Error(wrErr)
		return wrErr
	}

	// dogOwnerのSignUp
	resDogOwner, wrErr := ac.ah.SignUp(c, reqADOD)

	if wrErr != nil {
		return wrErr
	}

	// jwt処理
	return jwtProcessing(c, resDogOwner)
}

/*
jwt処理
*/
func jwtProcessing(c echo.Context, rdo dto.ResDogOwnerDto) error {
	logger := log.GetLogger(c).Sugar()

	// 秘密鍵取得
	secretKey := configs.FetchCondigStr("jwt.os.secret.key")
	jwtExpTime := configs.FetchCondigInt("jwt.exp.time")

	// jwt token生成
	signedToken, err := createToken(secretKey, uint64(rdo.DogOwnerID), jwtExpTime)
	if err != nil {
		logger.Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, success.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "dog owner successfully created",
		Token:   signedToken,
	})
}

/*
jwtのトークン生成
*/
func createToken(secretKey string, resAuthDogOwnerID uint64, expTime int) (string, error) {
	// JWTのペイロード
	claims := &dto.AccountClaims{
		ID: strconv.FormatUint(uint64(resAuthDogOwnerID), 10), // stringにコンバート
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expTime))), // 有効時間
		},
	}

	// token生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// tokenに署名
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

/*
GrantTypeの作成
*/
func createGrantType(c echo.Context, ih types.GrantType, pgh types.GrantType) (types.GrantType, error) {
	logger := log.GetLogger(c).Sugar()

	// GrantTypeヘッダーを取得
	grantTypeHeader := c.Request().Header.Get(string(ih))

	// GrantTypeヘッダーのバリデーション
	if err := dto.ValidateGrantTypeHeader(grantTypeHeader, string(pgh)); err != nil {
		err = wrErrors.NewWRError(err, "ヘッダー情報が異なります。", wrErrors.NewDogrunClientErrorEType())
		return "", err
	}

	// GrantTypeに型変換
	grantType := types.GrantType(grantTypeHeader)
	logger.Infof("grantTypeHeader: %v, Type: %T", grantType, grantType)

	return grantType, nil
}

// func (ac *authController) LogIn(c echo.Context) error {
// 	logger := log.GetLogger(c).Sugar()
// 	var reqADOD dto.ReqAuthDogOwnerDto = dto.ReqAuthDogOwnerDto{}

// 	if err := c.Bind(&reqADOD); err != nil {
// 		logger.Error(err)
// 		return c.JSON(http.StatusBadRequest, wrErrors.ErrorResponse{
// 			Code:    http.StatusBadRequest,
// 			Message: "Invalid format",
// 		})
// 	}
// 	logger.Infof("Request AuthDogOwner info: %v", reqADOD)

// 	// LogIn処理
// 	resAuthDogOwner, err := ac.ah.LogIn(c, reqADOD)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, wrErrors.ErrorResponse{
// 			Code:    http.StatusBadRequest,
// 			Message: "Invalid Request",
// 		})
// 	}

// 	// 秘密鍵取得
// 	secretKey := configs.FetchCondigStr("os.secret.key")

// 	// jwt token生成
// 	signedToken, err := createToken(secretKey, resAuthDogOwner.DogOwnerID, 72)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, wrErrors.ErrorResponse{
// 			Code:    http.StatusInternalServerError,
// 			Message: "Failed to sign token",
// 		})
// 	}

// 	return c.JSON(http.StatusCreated, success.SuccessResponse{
// 		Code:    http.StatusOK,
// 		Message: "Successful login",
// 		Token:   signedToken,
// 	})
// }

func (ac *authController) LogOut(c echo.Context) error { return nil }
