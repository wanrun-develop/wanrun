package model

import "time"

type AuthDogOwner struct {
	AuthDogOwnerID uint      `json:"authDogOwnerId" gorm:"primaryKey;column:auth_dog_owner_id;autoIncrement"`
	Password       string    `json:"password" gorm:"size:256;column:password;not null"`
	GrantType      int       `json:"grantType" gorm:"column:grant_type"`
	LoginAt        time.Time `json:"loginAt" gorm:"column:login_at;not null;autoCreateTime"`
	DogOwner       DogOwner  `json:"dogOwner" gorm:"foreignKey:DogOwnerID;references:DogOwnerID"`
	DogOwnerID     uint      `json:"DogOwnerId" gorm:"column:dog_owner_id;not null"`
}
