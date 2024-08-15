package model

import "time"

type DogOwner struct {
	DogOwnerID uint      `gorm:"primaryKey;column:dog_owner_id;autoIncrement"`
	Name       string    `gorm:"size:128;column:name;not null"`
	Email      string    `gorm:"size:255;column:email;not null"`
	Image      string    `gorm:"type:text;column:image"`
	Sex        string    `gorm:"size:1;column:sex"`
	CreateAt   time.Time `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:upd_at;not null;autoCreateTime"`
}

type DogOwnerRes struct {
	DogOwnerID uint      `gorm:"primaryKey;column:dog_owner_id;autoIncrement"`
	Name       string    `gorm:"size:128;column:name;not null"`
	Email      string    `gorm:"size:255;column:email;not null"`
	Image      string    `gorm:"type:text;column:image"`
	Sex        string    `gorm:"size:1;column:sex"`
	CreateAt   time.Time `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:upd_at;not null;autoCreateTime"`
}
