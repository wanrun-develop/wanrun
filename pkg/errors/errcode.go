package errors

import "fmt"

const (
	AUTH     int = 1
	DOG      int = 2
	DOGOWNER int = 3
	DOGRUN   int = 4
)

const (
	CLIENT int = 1
	SERVER int = 2
)

type errorContext struct {
	service   int
	errorType int
}

/*
認証機能のクライアントエラー
*/
func NewAuthClientErrorType() errorContext {
	return errorContext{AUTH, CLIENT}
}

/*
認証機能のサーバーエラー
*/
func NewAuthSeverErrorType() errorContext {
	return errorContext{AUTH, SERVER}
}

/*
ドッグ機能のクライアントエラー
*/
func NewDogClientErrorType() errorContext {
	return errorContext{DOG, CLIENT}
}

/*
ドッグ機能のサーバーエラー
*/
func NewDogServerErrorType() errorContext {
	return errorContext{DOG, SERVER}
}

/*
ドッグオーナー機能のクライアントエラー
*/
func NewDogownerClientErrorType() errorContext {
	return errorContext{DOGOWNER, CLIENT}
}

/*
ドッグオーナー機能のサーバーエラー
*/
func NewDogownerServerErrorType() errorContext {
	return errorContext{DOGOWNER, SERVER}
}

/*
ドッグラン機能のクライアントエラー
*/
func NewDogrunClientErrorType() errorContext {
	return errorContext{DOGRUN, CLIENT}
}

/*
ドッグラン機能のサーバーエラー
*/
func NewDogrunServerErrorType() errorContext {
	return errorContext{DOGRUN, SERVER}
}

func (t errorContext) String() string {
	// カスタムフォーマットで文字列を返す
	return fmt.Sprintf("%d-%d", t.service, t.errorType)
}
