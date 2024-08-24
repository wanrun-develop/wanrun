package model

import "time"

type DogOwner struct {
	DogOwnerID uint      `json:"dogOwnerId" gorm:"primaryKey;column:dog_owner_id;autoIncrement"`
	Name       string    `json:"name" gorm:"size:128;column:name;not null"`
	Email      string    `json:"email" gorm:"size:255;column:email;not null"`
	Image      string    `json:"image" gorm:"type:text;column:image"`
	Sex        string    `json:"sex" gorm:"size:1;column:sex"`
	CreateAt   time.Time `json:"createAt" gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt   time.Time `json:"updateAt" gorm:"column:upd_at;not null;autoCreateTime"`
}

type DogOwnerRes struct {
	DogOwnerID uint      `json:"dogOwnerId" gorm:"primaryKey;column:dog_owner_id;autoIncrement"`
	Name       string    `json:"name" gorm:"size:128;column:name;not null"`
	Email      string    `json:"email" gorm:"size:255;column:email;not null"`
	Image      string    `json:"image" gorm:"type:text;column:image"`
	Sex        string    `json:"sex" gorm:"size:1;column:sex"`
	CreateAt   time.Time `json:"createAt" gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt   time.Time `json:"updateAt" gorm:"column:upd_at;not null;autoCreateTime"`
}
