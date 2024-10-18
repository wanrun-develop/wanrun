package model

import (
	"database/sql"
	"time"

	"github.com/wanrun-develop/wanrun/internal/models/types"
	"github.com/wanrun-develop/wanrun/pkg/util"
)

type AuthDogOwner struct {
	AuthDogOwnerID         sql.NullInt64   `json:"authDogOwnerId" gorm:"primaryKey;column:auth_dog_owner_id;autoIncrement"`
	GrantType              types.GrantType `gorm:"column:grant_type"`
	AccessToken            sql.NullString  `gorm:"size:512;column:access_token"`
	RefreshToken           sql.NullString  `gorm:"size:512;column:refresh_token"`
	AccessTokenExpiration  util.CustomTime `gorm:"column:access_token_expiration"`
	RefreshTokenExpiration util.CustomTime `gorm:"column:refresh_token_expiration"`
	LoginAt                time.Time       `json:"loginAt" gorm:"column:login_at;not null;autoCreateTime"`

	DogOwner   DogOwner      `json:"dogOwner" gorm:"foreignKey:DogOwnerID;references:DogOwnerID"`
	DogOwnerID sql.NullInt64 `json:"DogOwnerId" gorm:"column:dog_owner_id;not null"`
}

type DogOwnerCredential struct {
	CredentialID sql.NullInt64 `gorm:"primaryKey;column:credential_id;autoIncrement"`
	// ProviderName   sql.NullString `gorm:"size:50;column:provider_name"`
	ProviderUserID sql.NullString `gorm:"size:256;column:provider_user_id"`
	Email          sql.NullString `gorm:"size:256;column:email"`
	PhoneNumber    sql.NullString `gorm:"size:15;column:phone_number"`
	Password       sql.NullString `gorm:"size:256;column:password"`
	LoginAt        sql.NullTime   `gorm:"column:login_at;autoCreateTime"`

	AuthDogOwner   AuthDogOwner  `gorm:"foreignKey:AuthDogOwnerID;references:AuthDogOwnerID"`
	AuthDogOwnerID sql.NullInt64 `gorm:"column:auth_dog_owner_id;not null"`
}
