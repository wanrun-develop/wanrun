package googleplace

import "strings"

// field mask
const (
	//IDs Only
	F_ID_IO     = "id"     //place id
	F_PHOTOS_IO = "photos" //画像
	//Location Only
	F_ADDRESSCOMPONENTS_LO     = "addressComponents"     //住所(構造型)
	F_ADRFORMATADDRESS_LO      = "adrFormatAddress"      //住所(ADRフォーマット)
	F_FORMATTEDADDRESS_LO      = "formattedAddress"      //住所(フォーマット)
	F_SHORTFORMATTEDADDRESS_LO = "shortFormattedAddress" //住所(区市町村以降)
	F_LOCATION_LO              = "location"              //経度/緯度
	F_VIEWPORT_LO              = "viewport"              //矩形（バウンディングボックス）
	F_PLUSCODE_LO              = "plusCode"              //場所コード
	F_TYPES_LO                 = "types"                 //place type
	//Basic
	F_DISPLAYNAME = "displayName" // 表示名
	F_RATING      = "rating"      //評価
)

type IFieldMask interface {
	getValue() string
}

// 基礎情報
type BaseField struct{}

func (b BaseField) getValue() string {
	return strings.Join([]string{F_ID_IO, F_SHORTFORMATTEDADDRESS_LO, F_LOCATION_LO}, ",")
}
