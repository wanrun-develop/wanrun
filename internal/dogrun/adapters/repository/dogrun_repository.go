package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/dogrun/core/dto"
	model "github.com/wanrun-develop/wanrun/internal/models"
	"github.com/wanrun-develop/wanrun/pkg/errors"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"gorm.io/gorm"
)

type IDogrunRepository interface {
	GetDogrunByPlaceID(c echo.Context, placeID string) (model.Dogrun, error)
	GetDogrunByID(id string) (model.Dogrun, error)
	GetDogrunByRectanglePointer(c echo.Context, condition dto.SearchAroudRectangleCondition) ([]model.Dogrun, error)
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
	if err := drr.db.Preload("DogrunTags.TagMst").Where("place_id = ?", placeID).Find(&dogrun).Error; err != nil {
		logger.Error(err)
		return model.Dogrun{}, errors.NewWRError(err, "DBからのデータ取得に失敗", errors.NewDogrunServerErrorEType())
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

/*
リクエストのボディの条件に基づいて、指定範囲内のドッグランを取得する
*/
func (drr *dogrunRepository) GetDogrunByRectanglePointer(c echo.Context, condition dto.SearchAroudRectangleCondition) ([]model.Dogrun, error) {
	logger := log.GetLogger(c).Sugar()
	dogruns := []model.Dogrun{}
	if err := drr.db.Preload("DogrunTags.TagMst").
		Where("longitude BETWEEN ? AND ?", condition.Target.Southwest.Longitude, condition.Target.Northeast.Longitude).
		Where("latitude BETWEEN ? AND ?", condition.Target.Southwest.Latitude, condition.Target.Northeast.Latitude).
		Find(&dogruns).Error; err != nil {
		logger.Error(err)
		return nil, errors.NewWRError(err, "DBからのデータ取得に失敗", errors.NewDogrunServerErrorEType())
	}
	return dogruns, nil
}
