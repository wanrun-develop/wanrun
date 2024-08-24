package repository

import (
	"github.com/labstack/echo/v4"
	model "github.com/wanrun-develop/wanrun/internal/models"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"gorm.io/gorm"
)

type IDogrunRepository interface {
	GetDogrunByPlaceID(c echo.Context, placeID string) (model.Dogrun, error)
	GetDogrunByID(id string) (model.Dogrun, error)
}

type dogrunRepository struct {
	db *gorm.DB
}

func NewDogrunRepository(db *gorm.DB) IDogrunRepository {
	return &dogrunRepository{db}
}

/*
PlaceIDで、ドッグランの取得
*/
func (drr *dogrunRepository) GetDogrunByPlaceID(c echo.Context, placeID string) (model.Dogrun, error) {
	logger := log.GetLogger(c).Sugar()
	dogrun := model.Dogrun{}
	if err := drr.db.Where("place_id = ?", placeID).Find(&dogrun).Error; err != nil {
		logger.Error(err)
		return dogrun, err
	}
	return dogrun, nil
}

/*
DogrunIDで、ドッグランの取得
*/
func (drr *dogrunRepository) GetDogrunByID(id string) (model.Dogrun, error) {
	dogrun := model.Dogrun{}
	if err := drr.db.Where("dogrun_id = ?", id).Find(&dogrun).Error; err != nil {
		return dogrun, err
	}
	return dogrun, nil
}
