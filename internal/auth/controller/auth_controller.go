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
	model "github.com/wanrun-develop/wanrun/internal/models"
	"github.com/wanrun-develop/wanrun/pkg/errors"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"github.com/wanrun-develop/wanrun/pkg/success"
)

type IAuthController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

type authController struct {
	ah handler.IAuthHandler
}

func NewAuthController(ah handler.IAuthHandler) IAuthController {
	return &authController{ah}
}

func (ac *authController) SignUp(c echo.Context) error {
	logger := log.GetLogger(c).Sugar()

	reqADOD := dto.ReqAuthDogOwnerDto{}
	if err := c.Bind(&reqADOD); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, errors.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid format",
		})
	}

	// dogOwnerのSignUp
	resAuthDogOwner, err := ac.ah.SignUp(reqADOD)
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to register dog owner information",
		})
	}

	// 秘密鍵取得
	secretKey := configs.FetchCondigStr("os.secret.key")

	// jwt token生成
	signedToken, err := createToken(secretKey, resAuthDogOwner.DogOwnerID, 72)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to sign token",
		})
	}
	return c.JSON(http.StatusCreated, success.SuccessResponse{
		Code:    http.StatusCreated,
		Message: "DogOwner created!!!",
		Token:   signedToken,
	})
}

/*
jwtのトークン生成
*/
func createToken(secretKey string, resAuthDogOwnerID uint, expTime int) (string, error) {
	// JWTのペイロード
	claims := &model.AccountClaims{
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

func (ac *authController) LogIn(c echo.Context) error  { return nil }
func (ac *authController) LogOut(c echo.Context) error { return nil }
