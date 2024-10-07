package dto

import "time"

// ドッグラン詳細画面での表示情報
type DogrunDetailDto struct {
	DogrunID        int            `json:"dogrun_id,omitempty"`
	DogrunManagerID int            `json:"dogrun_manager_id,omitempty"`
	PlaceId         string         `json:"place_id,omitempty"`
	Name            string         `json:"name"`
	Address         Address        `json:"address"`
	Location        Location       `json:"location"`
	BusinessStatus  string         `json:"business_status"`
	NowOpen         bool           `json:"now_open"`
	BusinessHour    BusinessHour   `json:"business_hour"`
	Description     string         `json:"description,omitempty"`
	GoogleRating    float32        `json:"google_rating,omitempty"`
	UserRatingCount int            `json:"user_rating_count,omitempty"`
	DogrunTags      []DogrunTagDto `json:"dogrun_tags,omitempty"`
	CreateAt        *time.Time     `json:"create_at,omitempty"`
	UpdateAt        *time.Time     `json:"update_at,omitempty"`
}

// ドッグラン一覧での表示情報
type DogrunListDto struct {
	DogrunID        int            `json:"dogrun_id,omitempty"`
	PlaceId         string         `json:"place_id,omitempty"`
	Name            string         `json:"name"`
	Address         Address        `json:"address"`
	Location        Location       `json:"location"`
	NowOpen         bool           `json:"now_open"`
	OpenTime        string         `json:"open_time"`
	CloseTime       string         `json:"close_time"`
	Description     string         `json:"description,omitempty"`
	GoogleRating    float32        `json:"google_rating,omitempty"`
	UserRatingCount int            `json:"user_rating_count,omitempty"`
	DogrunTags      []DogrunTagDto `json:"dogrun_tags,omitempty"`
	CreateAt        *time.Time     `json:"create_at,omitempty"`
	UpdateAt        *time.Time     `json:"update_at,omitempty"`
}

// 営業日情報
type BusinessHour struct {
	Regular RegularBusinessHour   `json:"regular"`
	Special []SpecialBusinessHour `json:"special,omitempty"`
}

// 通常営業日情報
type RegularBusinessHour struct {
	Sunday    DayBusinessTime `json:"0,omitempty"`
	Monday    DayBusinessTime `json:"1,omitempty"`
	Tuesday   DayBusinessTime `json:"2,omitempty"`
	Wednesday DayBusinessTime `json:"3,omitempty"`
	Thursday  DayBusinessTime `json:"4,omitempty"`
	Friday    DayBusinessTime `json:"5,omitempty"`
	Saturday  DayBusinessTime `json:"6,omitempty"`
}

type DayBusinessTime struct {
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
	IsAllDay  bool   `json:"is_all_day"`
	IsHoliday bool   `json:"is_holiday"`
}

// 特別営業日情報
type SpecialBusinessHour struct {
	Date string `json:"date"`
	DayBusinessTime
}

// ドッグランタグ情報
type DogrunTagDto struct {
	DogrunTagId int    `json:"dogrun_tag_id"`
	TagId       int    `json:"tag_id"`
	TagName     string `json:"tag_idag_name"`
	Description string `json:"description"`
}

// 軽度・緯度情報
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// 住所情報
type Address struct {
	PostCode string `json:"postcode"`
	Address  string `json:"address"`
}
