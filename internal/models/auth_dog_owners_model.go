package model

type AuthDogOwner struct {
	AuthDogOwnerID uint   `json:"authDogOwnerId" gorm:"primaryKey;column:auth_dog_owner_id;autoIncrement"`
	Password       string `json:"password" gorm:"size:256;column:password;not null"`
	GrantType      int    `json:"grantType" gorm:"column:grant_type"`
	// PassRegistAt   time.Time `json:"passRegistAt" gorm:"column:pass_regist_at;not null;autoCreateTime"`
	DogOwner   DogOwner `json:"dogOwner" gorm:"foreignKey:DogOwnerID;references:DogOwnerID"`
	DogOwnerID uint     `json:"DogOwnerId" gorm:"column:dog_owner_id;not null"`
}

type ResAuthDogOwner struct {
	AuthDogOwnerID uint   `json:"authDogOwnerId"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Sex            string `json:"sex"`
}
