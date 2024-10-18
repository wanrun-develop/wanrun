/*
Enum型用のカスタム型の定義
*/

package types

import "fmt"

type GrantType string

// フロントエンドからのGrantTypeの値
const (
	OAUTH_IDENTIFICATION_HEADER GrantType = "GrantType" // 識別用のキー
	OAUTH_GRANT_TYPE_HEADER     GrantType = "OAUTH"
	PASSWORD_GRANT_TYPE_HEADER  GrantType = "PASSWORD"
)

func AuthenticateUser(grantType GrantType) (string, error) {
	switch grantType {
	case OAUTH_IDENTIFICATION_HEADER:
		return string(OAUTH_IDENTIFICATION_HEADER), nil
	case OAUTH_GRANT_TYPE_HEADER:
		return string(OAUTH_GRANT_TYPE_HEADER), nil
	case PASSWORD_GRANT_TYPE_HEADER:
		return string(PASSWORD_GRANT_TYPE_HEADER), nil
	default:
		return "", fmt.Errorf("unknown authentication type")
	}
}
