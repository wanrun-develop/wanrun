package core

// token authentication middleware用の定数

const (
	CONTEXT_KEY   string = "user_info"
	TOKEN_LOOK_UP string = "header:Authorization:Bearer " // `Bearer `しか切り取れないのでスペースが多い場合は未対応
	// Cookie認証用の定数
	COOKIE_TOKEN_LOOK_UP string = "cookie:" + AUTH_COOKIE_NAME
)

// role
const (
	SYSTEM              int = 0
	DOGRUNMG_ROLE       int = 1
	DOGRUNMG_ADMIN_ROLE int = 2
	DOGOWNER_ROLE       int = 3
	GENERAL             int = 100
)

// 一般ユーザーのUserID
const (
	GENERAL_USER_ID     int64  = -999
	GENERAL_USER_JWT_ID string = "general"
)

// 認証方法
const (
	PASSWORD string = "password"
	REFRESH  string = "refresh"
)

// Cookie設定
const (
	AUTH_COOKIE_NAME      string = "wanrun_auth_token"
	AUTH_COOKIE_PATH      string = "/"
	AUTH_COOKIE_MAX_AGE   int    = 24 * 60 * 60 // 24時間（秒単位）
	AUTH_COOKIE_HTTP_ONLY bool   = true
	AUTH_COOKIE_SECURE    bool   = true
	AUTH_COOKIE_SAME_SITE string = "strict"
)
