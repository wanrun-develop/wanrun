/*
Enum型用のカスタム型の定義
*/

package types

type GrantType string

// フロントエンドからのGrantTypeの値
const (
	OAUTH_IDENTIFICATION_HEADER GrantType = "GrantType" // 識別用のキー
	OAUTH_GRANT_TYPE_HEADER     GrantType = "OAUTH"
	PASSWORD_GRANT_TYPE_HEADER  GrantType = "PASSWORD"
)
