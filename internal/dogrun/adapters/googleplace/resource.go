package googleplace

import "fmt"

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type SearchTextBaseResource struct {
	Places        []BaseResource `json:"places"`
	NextPageToken *string        `json:"nextPageToken"`
}

type BaseResource struct {
	ID                    string             `json:"id"`
	Location              Location           `json:"location"`
	ShortFormattedAddress string             `json:"shortFormattedAddress"`
	AddressComponents     []AddressComponent `json:"addressComponents"`
	DisplayName           LocalizedText      `json:"displayName"`
	Rating                float32            `json:"rating"`
	UserRatingCount       int                `json:"userRatingCount"`
	BusinessStatus        string             `json:"businessStatus"`
	OpeningHours          OpeningHours       `json:"regularOpeningHours"`
	Summary               LocalizedText      `json:"editorialSummary"`
	Photos                []PhotoObject      `json:"photos"`
}

/*
BaseResourceが空かの判定
*/
func (r *BaseResource) IsEmpty() bool {
	return r.ID == ""
}

/*
BaseResourceが空でないかの判定
*/
func (r *BaseResource) IsNotEmpty() bool {
	return !r.IsEmpty()
}

/*
OpeningHoursが空かの判定
*/
func (o *OpeningHours) IsEmpty() bool {
	return len(o.Periods) == 0 && len(o.WeekdayDescriptions) == 0
}

/*
OpeningHoursが空でないかの判定
*/
func (o *OpeningHours) IsNotEmpty() bool {
	return !o.IsEmpty()
}

type LocalizedText struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
}

// google写真情報
type PhotoObject struct {
	Name     string `json:"name"`
	WidthPx  uint   `json:"widthPx"`
	HeightPx uint   `json:"heightPx"`
}

// 構造型住所
type AddressComponent struct {
	LongText  string   `json:"longText"`
	ShortText string   `json:"shortText"`
	Types     []string `json:"types"`
}

const (
	ADDRESSCOMPONENT_TYPES_POSTAL_CODE string = "postal_code" //addressComponents.typeの郵便番号
)

// 営業時間
type OpeningHours struct {
	OpenNow             bool                 `json:"openNow"`
	Periods             []OpeningHoursPeriod `json:"periods"`
	WeekdayDescriptions []string             `json:"weekdayDescriptions"`
}

/*
曜日(数値)より、対象の営業開始/営業終了時間を返す
*/
func (oh *OpeningHours) FetchTargetPeriod(day int) (*OpeningHoursPeriodInfo, *OpeningHoursPeriodInfo) {
	var targetOpenPeriod *OpeningHoursPeriodInfo
	var targetClosePeriod *OpeningHoursPeriodInfo

	for _, v := range oh.Periods {
		if v.Open.Day == day {
			targetOpenPeriod = &v.Open
		}
		if v.Close.Day == day {
			targetClosePeriod = &v.Close
		}
	}

	return targetOpenPeriod, targetClosePeriod
}

// 営業時間 period
type OpeningHoursPeriod struct {
	Open  OpeningHoursPeriodInfo `json:"open"`
	Close OpeningHoursPeriodInfo `json:"close"`
}

// 営業時間 period info
type OpeningHoursPeriodInfo struct {
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

/*
OpeningHoursPeriodInfoの時間をHH:mm:ss(文字列)で返す。
1桁の場合は頭に0を付与する
*/
func (o *OpeningHoursPeriodInfo) FormatTime() string {
	// 1桁の場合は頭に0をつける
	hh := fmt.Sprintf("%02d", o.Hour)
	mm := fmt.Sprintf("%02d", o.Minute)

	// "HH:mm" 形式で返す
	return fmt.Sprintf("%s:%s:00", hh, mm)
}

// google places photo mediaのレスポンス
type PhotoMediaResource struct {
	Name     string `json:"name"`
	PhotoUri string `json:"photoUri"`
}
