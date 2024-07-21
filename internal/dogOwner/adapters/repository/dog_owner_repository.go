package repository

import (
	"github.com/wanrun-develop/wanrun/internal/dogOwner/core/models"
	"gorm.io/gorm"
)

/*
dog_ownerのrepositoryの作成
*/
type DogOwnerRepository struct {
	dbConn *gorm.DB
}

/*
repositoryの作成
*/
func NewDogOwnerRepository(dbConn *gorm.DB) *DogOwnerRepository {
	return &DogOwnerRepository{dbConn: dbConn}
}

/*
引数のmodels.DogOwnerをinsertする
*/
func (r *DogOwnerRepository) InsertDogOwner(dogOwner *models.DogOwner) {
	r.dbConn.Create(&dogOwner)
}
