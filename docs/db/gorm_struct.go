package temp

import (
	"time"
)

type DogOwner struct {
	DogOwnerID int       `gorm:"primaryKey;column:dog_owner_id;autoIncrement"`
	Name       string    `gorm:"size:128;column:name;not null"`
	Email      string    `gorm:"size:255;column:email;not null"`
	Image      string    `gorm:"type:text;column:image"`
	Sex        string    `gorm:"size:1;column:sex"`
	CreateAt   time.Time `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:upd_at;not null;autoCreateTime"`
}

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

type DogTypeMst struct {
	DogTypeID int    `gorm:"primaryKey;column:dog_type_id;autoIncrement"`
	Name      string `gorm:"size:64;column:name;not null"`
}

type InjectionCertification struct {
	InjectionCertificationID int       `gorm:"primaryKey;column:injection_certification_id;autoIncrement"`
	DogID                    int       `gorm:"column:dog_id;not null;foreignKey:DogID"`
	Type                     int       `gorm:"column:type;not null"`
	File                     string    `gorm:"type:text;column:file;not null"`
	CreateAt                 time.Time `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt                 time.Time `gorm:"column:upd_at;not null;autoCreateTime"`
}

type DogrunManager struct {
	DogrunManagerID int       `gorm:"primaryKey;column:dogrun_manager_id;autoIncrement"`
	Name            string    `gorm:"size:128;column:name"`
	Email           string    `gorm:"size:255;column:email;not null"`
	CreateAt        time.Time `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt        time.Time `gorm:"column:upd_at;not null;autoCreateTime"`
}

type Dogrun struct {
	DogrunID        int       `gorm:"primaryKey;column:dogrun_id;autoIncrement"`
	DogrunManagerID int       `gorm:"column:dogrun_manager_id;foreignKey:DogrunManagerID"`
	Name            string    `gorm:"size:256;column:name;not null"`
	Address         string    `gorm:"size:256;column:address"`
	PostCode        string    `gorm:"size:8;column:postcode"`
	BusinessDay     int       `gorm:"column:business_day"`
	Holiday         int       `gorm:"column:holiday"`
	OpenTime        time.Time `gorm:"column:open_time"`
	CloseTime       time.Time `gorm:"column:close_time"`
	Description     string    `gorm:"type:text;column:description"`
	CreateAt        time.Time `gorm:"column:reg_at;not null;autoCreateTime"`
	UpdateAt        time.Time `gorm:"column:upd_at;not null;autoCreateTime"`
}

type DogrunImage struct {
	DogrunImageID int       `gorm:"primaryKey;column:dogrun_image_id;autoIncrement"`
	DogrunID      int       `gorm:"column:dogrun_id;not null;foreignKey:DogrunID"`
	Image         string    `gorm:"type:text;column:image;not null"`
	SortOrder     int       `gorm:"column:sort_order"`
	UploadAt      time.Time `gorm:"column:upload_at"`
}

type DogrunTag struct {
	DogrunTagID int `gorm:"primaryKey;column:dogrun_tag_id;autoIncrement"`
	DogrunID    int `gorm:"column:dogrun_id;not null;foreignKey:DogrunID"`
	TagID       int `gorm:"column:tag_id;not null;foreignKey:TagID"`
}

type TagMst struct {
	TagID       int    `gorm:"primaryKey;column:tag_id;autoIncrement"`
	TagName     string `gorm:"size:64;column:tag_name;not null"`
	Description string `gorm:"type:text;column:description"`
}

type AuthDogOwner struct {
	AuthDogOwnerID int       `gorm:"primaryKey;column:auth_dog_owner_id;autoIncrement"`
	DogOwnerID     int       `gorm:"column:dog_owner_id;not null;foreignKey:DogOwnerID"`
	Password       string    `gorm:"size:256;column:password"`
	GrantType      int       `gorm:"column:grant_type;not null"`
	LoginAt        time.Time `gorm:"column:login_at"`
}

type AuthDogrunManager struct {
	AuthDogrunManagerID int       `gorm:"primaryKey;column:auth_dogrun_manager_id;autoIncrement"`
	DogrunManagerID     int       `gorm:"column:dogrun_manager_id;not null;foreignKey:DogrunManagerID"`
	Password            string    `gorm:"size:256;column:password"`
	GrantType           int       `gorm:"column:grant_type;not null"`
	LoginAt             time.Time `gorm:"column:login_at"`
}
