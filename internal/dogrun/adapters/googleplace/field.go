package googleplace

import "strings"

// field mask
const (
	//IDs Only
	f_id_io     = "id"     //place id
	f_photos_io = "photos" //画像
	//Location Only
	f_addressComponents_lo     = "addressComponents"     //住所(構造型)
	f_adrFormatAddress_lo      = "adrFormatAddress"      //住所(ADRフォーマット)
	f_formattedAddress_lo      = "formattedAddress"      //住所(フォーマット)
	f_shortFormattedAddress_lo = "shortFormattedAddress" //住所(区市町村以降)
	f_location_lo              = "location"              //経度/緯度
	f_viewport_lo              = "viewport"              //矩形（バウンディングボックス）
	f_plusCode_lo              = "plusCode"              //場所コード
	f_types_lo                 = "types"                 //place type
	//Basic
	f_displayName = "displayName" // 表示名
	f_rating      = "rating"      //評価
)

type IFieldMask interface {
	getValue() string
}

// 基礎情報
type BaseField struct{}

func (b BaseField) getValue() string {
	return strings.Join([]string{f_id_io, f_shortFormattedAddress_lo, f_location_lo}, ",")
}
