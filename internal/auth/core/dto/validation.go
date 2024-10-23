package dto

import (
	"errors"

	"github.com/wanrun-develop/wanrun/internal/models/types"
	wrErrors "github.com/wanrun-develop/wanrun/pkg/errors"
)

var (
	// Header
	errInvalidHeader     = errors.New("request header is wrong")
	errEmptyHeader       = errors.New("grant type request header is empty")
	errUnsupportedHeader = errors.New("unsupported grant type")

	// OAuth
	errOAuthFailed     = errors.New("oauth authorization failed")
	errOAuthInvalidReq = errors.New("neither code nor error parameter found")

	// Password
	errPasswordBothEmpty  = errors.New("email and phoneNumber cannot be both empty")
	errPasswordBothFilled = errors.New("both email and phoneNumber cannot have values")
)

/*
GrantTypeのheaderバリデーション
*/
func ValidateGrantTypeHeader(gotGrantTypeHeader string, correctGrantType string) error {
	// 空の時のエラーチェック
	if gotGrantTypeHeader == "" {
		wrErr := wrErrors.NewWRError(
			errEmptyHeader,
			"ヘッダーが空です",
			wrErrors.NewDogownerClientErrorEType(),
		)
		return wrErr
	}

	// "PASSWORD"か"OAUTH"以外が来ないかチェック
	if gotGrantTypeHeader != string(types.OAUTH_GRANT_TYPE_HEADER) && gotGrantTypeHeader != string(types.PASSWORD_GRANT_TYPE_HEADER) {
		wrErr := wrErrors.NewWRError(
			errUnsupportedHeader,
			"PASSWORDかOAUTHヘッダー以外がリクエストされております",
			wrErrors.NewDogownerClientErrorEType(),
		)
		return wrErr
	}

	// 取得ヘッダーと正しいヘッダーのチェック
	if gotGrantTypeHeader != correctGrantType {
		wrErr := wrErrors.NewWRError(
			errInvalidHeader,
			"ヘッダー情報が違います",
			wrErrors.NewDogownerClientErrorEType(),
		)
		return wrErr
	}

	// それ以外はOK
	return nil
}

/*
OAuthのクエリパラメータのバリデーション
*/
func ValidateOAuthResCode(authorizationCode string, oauthErrorCode string) error {
	// "error" パラメータがある場合はエラーレスポンスを返す
	if oauthErrorCode != "" {
		wrErr := wrErrors.NewWRError(
			errOAuthFailed,
			"認証に失敗しました。",
			wrErrors.NewDogownerClientErrorEType(),
		)
		return wrErr
	}

	// "code" パラメータがある場合はそのまま処理
	if authorizationCode != "" {
		return nil
	}

	// どちらのパラメータもない場合は不正なリクエストとしてエラーを返す
	return errOAuthInvalidReq
}
