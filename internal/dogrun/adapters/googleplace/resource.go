package googleplace

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type BaseResource struct {
	ID                    string             `json:"id"`
	Location              Location           `json:"location"`
	ShortFormattedAddress string             `json:"shortFormattedAddress"`
	AddressComponents     []AddressComponent `json:"addressComponents"`
	DisplayName           LocalizedText      `json:"displayName"`
	Rating                float32            `json:"rating"`
	BusinessStatus        string             `json:"businessStatus"`
	OpeningHours          OpeningHours       `json:"regularOpeningHours"`
}

type LocalizedText struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
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
