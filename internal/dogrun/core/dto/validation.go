package dto

import "github.com/go-playground/validator/v10"

// 経度カスタムバリデーション
func VLatitude(fl validator.FieldLevel) bool {
	lat := fl.Field().Float()
	return lat >= -90 && lat <= 90
}

// 緯度カスタムバリデーション
func VLongitude(fl validator.FieldLevel) bool {
	lon := fl.Field().Float()
	return lon >= -180 && lon <= 180
}
