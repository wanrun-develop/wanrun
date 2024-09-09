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

type Type struct {
	Service   int
	ErrorType int
}

/*
認証機能のクライアントエラー
*/
func NewAuthClientErrorType() Type {
	return Type{AUTH, CLIENT}
}

/*
認証機能のサーバーエラー
*/
func NewAuthSeverErrorType() Type {
	return Type{AUTH, SERVER}
}

/*
ドッグ機能のクライアントエラー
*/
func NewDogClientErrorType() Type {
	return Type{DOG, CLIENT}
}

/*
ドッグ機能のサーバーエラー
*/
func NewDogServerErrorType() Type {
	return Type{DOG, SERVER}
}

/*
ドッグオーナー機能のクライアントエラー
*/
func NewDogownerClientErrorType() Type {
	return Type{DOGOWNER, CLIENT}
}

/*
ドッグオーナー機能のサーバーエラー
*/
func NewDogownerServerErrorType() Type {
	return Type{DOGOWNER, SERVER}
}

/*
ドッグラン機能のクライアントエラー
*/
func NewDogrunClientErrorType() Type {
	return Type{DOGRUN, CLIENT}
}

/*
ドッグラン機能のサーバーエラー
*/
func NewDogrunServerErrorType() Type {
	return Type{DOGRUN, SERVER}
}

func (t Type) String() string {
	// カスタムフォーマットで文字列を返す
	return fmt.Sprintf("%d-%d", t.Service, t.ErrorType)
}
