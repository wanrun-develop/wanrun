package googleplace

import (
	"fmt"
	"strings"
)

// field mask
const (
	//IDs Only
	F_ID_IO            = "id"            //place id
	F_PHOTOS_IO        = "photos"        //画像
	F_NEXTPAGETOKEN_IO = "nextPageToken" //次のページトークン(searchTextで使用)
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
	F_DISPLAYNAME_B    = "displayName"    // 表示名
	F_RATING_B         = "rating"         //評価
	F_BUSINESSSTATUS_B = "businessStatus" //評価(0~5)
	//Advanced
	F_USERRATINGCOUNT_A     = "userRatingCount"     //評価数
	F_REGULAROPENINGHOURS_A = "regularOpeningHours" //営業時間
	F_CURRENTOPENINGHOURS_A = "currentOpeningHours" //今日を含む７日間の営業日
	F_WEBSITEURI_A          = "websiteUri"          //webサイトURL

)

// リクエストに使うfieldMaskたち
var BASE_FILED_MASK = []string{
	F_ID_IO,
	F_SHORTFORMATTEDADDRESS_LO,
	F_ADDRESSCOMPONENTS_LO,
	F_LOCATION_LO,
	F_DISPLAYNAME_B,
	F_RATING_B,
	F_BUSINESSSTATUS_B,
	F_REGULAROPENINGHOURS_A,
	F_WEBSITEURI_A,
}

type IFieldMask interface {
	getValue() string
	getValueWPlaces() string
	getValueWPlacesAndNextPageToken() string
}

// 詳細情報用
type BaseField struct{}

func (b BaseField) getValue() string {
	return strings.Join(BASE_FILED_MASK, ",")
}

/*
search nearbyようにfieldに"palce."のプレフィックスを付与する
*/
func (b BaseField) getValueWPlaces() string {
	// BASE_FILED_MASK のコピーを作成
	fieldsWithPlace := make([]string, len(BASE_FILED_MASK))
	copy(fieldsWithPlace, BASE_FILED_MASK)

	// "place." をそれぞれの定数に付与
	for i, field := range fieldsWithPlace {
		fieldsWithPlace[i] = "places." + field
	}

	// カンマ区切りで連結
	return strings.Join(fieldsWithPlace, ",")
}

func (b BaseField) getValueWPlacesAndNextPageToken() string {
	basePlacesFilesMask := b.getValueWPlaces()
	return fmt.Sprintf("%s%s%s", basePlacesFilesMask, ",", F_NEXTPAGETOKEN_IO)
}
