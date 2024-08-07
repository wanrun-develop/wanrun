package model

import "time"

type Dog struct {
	DogID      int       `gorm:"primaryKey;column:dog_id;autoIncrement"`
	DogOwnerID int       `gorm:"column:dog_owner_id;not null;foreignKey:DogOwnerID"`
	Name       string    `gorm:"size:128;column:name;not null"`
	DogTypeID  int       `gorm:"column:dog_type_id"`
	Weight     int       `gorm:"column:weight"`
	Sex        string    `gorm:"size:1;column:sex"`
	Image      string    `gorm:"type:text;column:image"`
	CreateAt   time.Time `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:upd_at;not null;autoCreateTime"`
}

type DogRes struct {
	DogID      int       `gorm:"primaryKey;column:dog_id;autoIncrement"`
	DogOwnerID int       `gorm:"column:dog_owner_id;not null;foreignKey:DogOwnerID"`
	Name       string    `gorm:"size:128;column:name;not null"`
	DogTypeID  int       `gorm:"column:dog_type_id"`
	Weight     int       `gorm:"column:weight"`
	Sex        string    `gorm:"size:1;column:sex"`
	Image      string    `gorm:"type:text;column:image"`
	CreateAt   time.Time `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:upd_at;not null;autoCreateTime"`
}
