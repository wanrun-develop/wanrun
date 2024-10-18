package model

import (
	"database/sql"

	"github.com/wanrun-develop/wanrun/pkg/util"
)

const (
	EmailField       = "email"
	PhoneNumberField = "phone_number"
)

type Dogrun struct {
	DogrunID        sql.NullInt64   `gorm:"primaryKey;column:dogrun_id;autoIncrement"`
	DogrunManagerID sql.NullInt64   `gorm:"column:dogrun_manager_id;foreignKey:DogrunManagerID"`
	PlaceId         sql.NullString  `gorm:"size:256;column:place_id"`
	Name            sql.NullString  `gorm:"size:256;column:name"`
	Address         sql.NullString  `gorm:"size:256;column:address"`
	PostCode        sql.NullString  `gorm:"size:8;column:postcode"`
	BusinessDay     sql.NullInt64   `gorm:"column:business_day"`
	Holiday         sql.NullInt64   `gorm:"column:holiday"`
	OpenTime        util.CustomTime `gorm:"column:open_time"`
	CloseTime       util.CustomTime `gorm:"column:close_time"`
	Description     sql.NullString  `gorm:"type:text;column:description"`
	CreateAt        sql.NullTime    `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt        sql.NullTime    `gorm:"column:upd_at;not null;autoCreateTime"`

	//リレーション
	// DogrunImages []DogrunImage `gorm:"foreignKey:DogrunID"`
	DogrunTags []DogrunTag `gorm:"foreignKey:DogrunID;references:DogrunID"`
}

type DogrunTag struct {
	DogrunTagID sql.NullInt64 `gorm:"primaryKey;column:dogrun_tag_id;autoIncrement"`
	DogrunID    sql.NullInt64 `gorm:"column:dogrun_id;not null"`
	TagID       sql.NullInt64 `gorm:"column:tag_id;not null"`

	//リレーション
	// Dogrun Dogrun `gorm:"foreignKey:DogrunID"`
	TagMst TagMst `gorm:"foreignKey:TagID;references:TagID"`
}

type TagMst struct {
	TagID       sql.NullInt64  `gorm:"primaryKey;column:tag_id;autoIncrement"`
	TagName     sql.NullString `gorm:"size:64;column:tag_name;not null"`
	Description sql.NullString `gorm:"type:text;column:description"`
}

// GORMにテーブル名を指定
func (TagMst) TableName() string {
	return "tag_mst"
}

type DogrunImage struct {
	DogrunImageID int64         `gorm:"primaryKey;column:dogrun_image_id;autoIncrement"`
	DogrunID      int64         `gorm:"column:dogrun_id;not null"`
	Image         string        `gorm:"type:text;column:image;not null"`
	SortOrder     sql.NullInt64 `gorm:"column:sort_order"`
	UploadAt      sql.NullTime  `gorm:"column:upload_at"`

	//リレーション
	Dogrun Dogrun `gorm:"foreignKey:DogrunID"`
}
