package model

import "time"

type Dog struct {
	DogID      int       `json:"dogId" gorm:"primaryKey;column:dog_id;autoIncrement"`
	DogOwnerID int       `json:"dogOwnerId" gorm:"column:dog_owner_id;not null;foreignKey:DogOwnerID"`
	Name       string    `json:"name" gorm:"size:128;column:name;not null"`
	DogTypeID  int       `json:"dogTypeId" gorm:"column:dog_type_id"`
	Weight     int       `json:"weight" gorm:"column:weight"`
	Sex        string    `json:"sex" gorm:"size:1;column:sex"`
	Image      string    `json:"image" gorm:"type:text;column:image"`
	CreateAt   time.Time `json:"createAt" gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt   time.Time `json:"updateAt" gorm:"column:upd_at;not null;autoCreateTime"`
}

type DogRes struct {
	DogID      int       `json:"dogId" gorm:"primaryKey;column:dog_id;autoIncrement"`
	DogOwnerID int       `json:"dogOwnerId" gorm:"column:dog_owner_id;not null;foreignKey:DogOwnerID"`
	Name       string    `json:"name" gorm:"size:128;column:name;not null"`
	DogTypeID  int       `json:"dogTypeId" gorm:"column:dog_type_id"`
	Weight     int       `json:"weight" gorm:"column:weight"`
	Sex        string    `json:"sex" gorm:"size:1;column:sex"`
	Image      string    `json:"image" gorm:"type:text;column:image"`
	CreateAt   time.Time `json:"createAt" gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt   time.Time `json:"updateAt" gorm:"column:upd_at;not null;autoCreateTime"`
}
