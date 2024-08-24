package model

import (
	"database/sql"
	"time"
)

type Dogrun struct {
	DogrunID        sql.NullInt64  `gorm:"primaryKey;column:dogrun_id;autoIncrement"`
	DogrunManagerID sql.NullInt64  `gorm:"column:dogrun_manager_id;foreignKey:DogrunManagerID"`
	PlaceId         sql.NullString `gorm:"size:256;column:place_id"`
	Name            sql.NullString `gorm:"size:256;column:name"`
	Address         sql.NullString `gorm:"size:256;column:address"`
	PostCode        sql.NullString `gorm:"size:8;column:postcode"`
	BusinessDay     sql.NullInt64  `gorm:"column:business_day"`
	Holiday         sql.NullInt64  `gorm:"column:holiday"`
	OpenTime        CustomTime     `gorm:"column:open_time"`
	CloseTime       CustomTime     `gorm:"column:close_time"`
	Description     sql.NullString `gorm:"type:text;column:description"`
	CreateAt        sql.NullTime   `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt        sql.NullTime   `gorm:"column:upd_at;not null;autoCreateTime"`
}

type DogrunTag struct {
}

/*
時間型用の構造体
*/
type CustomTime struct {
	sql.NullTime
}

/*
時間のみの構造体に変換
gormはpsqlのtime型を自動でキャストできないようなので、実装
*/
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		ct.Valid = false
		return nil
	}
	ct.Valid = true
	var t time.Time
	switch v := value.(type) {
	case string:
		var err error
		t, err = time.Parse("15:04:05", v)
		if err != nil {
			return err
		}
	case time.Time:
		t = v
	default:
		ct.Valid = false
		return nil
	}
	ct.Time = t
	return nil
}
