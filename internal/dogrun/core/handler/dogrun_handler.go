package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/internal/dogrun/adapters/googleplace"
	"github.com/wanrun-develop/wanrun/internal/dogrun/adapters/repository"
	"github.com/wanrun-develop/wanrun/internal/dogrun/core/dto"
	model "github.com/wanrun-develop/wanrun/internal/models"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"github.com/wanrun-develop/wanrun/pkg/util"
)

const (
	SEARCH_TEXT_MAX_REQUEST_TIMES = 3 //searchTextの最大リクエスト数。pageSizeを20に指定すると、20*3=60個まで取得する
)

type IDogrunHandler interface {
	GetDogrunDetail(echo.Context, string) (*dto.DogrunDetailDto, error)
	GetDogrunByID(string)
	SearchAroundDogruns(echo.Context, dto.SearchAroudRectangleCondition) ([]googleplace.BaseResource, error)
}

type dogrunHandler struct {
	rest googleplace.IRest
	drr  repository.IDogrunRepository
}

func NewDogrunHandler(rest googleplace.IRest, drr repository.IDogrunRepository) IDogrunHandler {
	return &dogrunHandler{rest, drr}
}

func (h *dogrunHandler) GetDogrunDetail(c echo.Context, placeID string) (*dto.DogrunDetailDto, error) {
	logger := log.GetLogger(c).Sugar()
	//base情報のFieldを使用
	var baseFiled googleplace.IFieldMask = googleplace.BaseField{}
	//place情報の取得
	resG, err := h.rest.GETPlaceInfo(c, placeID, baseFiled)
	if err != nil {
		return nil, err
	}
	logger.Info("Google Place APIによって、ドッグラン情報の取得成功")

	// JSONデータを構造体にデコード
	var dogrunG googleplace.BaseResource
	err = json.Unmarshal(resG, &dogrunG)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.Info("Unmarshal成功")

	//dbから取得
	dogrunD, err := h.drr.GetDogrunByPlaceID(c, placeID)
	if err != nil {
		return nil, err
	}

	//情報選定
	resDogDetail := resolveDogrunDetail(dogrunG, dogrunD)
	return &resDogDetail, nil
}

/*
Google情報とDB情報から、ドッグラン詳細情報を作成
基本的に、DB情報をドッグランマネージャーからの手動更新がある前提で、優先情報とする
*/
func resolveDogrunDetail(dogrunG googleplace.BaseResource, dogrunD model.Dogrun) dto.DogrunDetailDto {

	// dogrunDetail := dto.DogrunDetailDto{}
	dogrunDetail := dto.DogrunDetailDto{
		DogrunID:        int(dogrunD.DogrunID.Int64),
		DogrunManagerID: int(dogrunD.DogrunManagerID.Int64),
		PlaceId:         dogrunG.ID,
		Name:            util.ChooseStringValidValue(dogrunD.Name, dogrunG.DisplayName.Text),
		Address:         resolveDogrunAddress(dogrunG, dogrunD),
		Location: dto.Location{
			Latitude:  dogrunG.Location.Latitude,
			Longitude: dogrunG.Location.Longitude,
		},
		BusinessStatus: dogrunG.BusinessStatus,
		NowOpen:        resolveBusinessStatus(dogrunD),
		BusinessDay:    int(dogrunD.BusinessDay.Int64), // TODO: 一旦、営業日時のテーブル設計再検討
		Holiday:        int(dogrunD.Holiday.Int64),     // TODO: 一旦、営業日時のテーブル設計再検討
		OpenTime:       resolveBuisnessTime(dogrunG.OpeningHours, dogrunD.OpenTime.NullTime, true),
		CloseTime:      resolveBuisnessTime(dogrunG.OpeningHours, dogrunD.CloseTime.NullTime, false),
		Description:    util.ChooseStringValidValue(dogrunD.Description, ""),
		DogrunTags:     resolveDogrunTagInfo(dogrunD), // ドッグランタグ情報
	}

	return dogrunDetail
}

/*
DBからドッグランタグ情報を取得
*/
func resolveDogrunTagInfo(dogrunD model.Dogrun) []dto.DogrunTagDto {
	dogrunTag := dogrunD.DogrunTags

	var dogrunTagInfos []dto.DogrunTagDto

	for _, v := range dogrunTag {
		dogrunTagInfo := dto.DogrunTagDto{
			DogrunTagId: int(v.DogrunID.Int64),
			TagId:       int(v.TagID.Int64),
			TagName:     v.TagMst.TagName.String,
			Description: v.TagMst.Description.String,
		}
		dogrunTagInfos = append(dogrunTagInfos, dogrunTagInfo)
	}
	return dogrunTagInfos
}

/*
住所情報の選定
*/
func resolveDogrunAddress(dogrunG googleplace.BaseResource, dogrunD model.Dogrun) dto.Address {

	address := util.ChooseStringValidValue(dogrunD.Address, dogrunG.ShortFormattedAddress)

	var gPostCodeComponent googleplace.AddressComponent
	for _, v := range dogrunG.AddressComponents {
		if slices.Contains(v.Types, googleplace.ADDRESSCOMPONENT_TYPES_POSTAL_CODE) {
			gPostCodeComponent = v
			break
		}
	}
	postCode := util.ChooseStringValidValue(dogrunD.PostCode, gPostCodeComponent.LongText)

	return dto.Address{PostCode: postCode, Address: address}
}

//TODO Google側の情報使うか検討中
/*
営業時間から、現在が営業中かを判定
*/
func resolveBusinessStatus(dogrunD model.Dogrun) bool {
	openTime := dogrunD.OpenTime.Time
	closeTime := dogrunD.CloseTime.Time

	now := time.Now()

	// 時間部分だけを取り出す
	openTime = time.Date(now.Year(), now.Month(), now.Day(), openTime.Hour(), openTime.Minute(), openTime.Second(), 0, time.UTC)
	closeTime = time.Date(now.Year(), now.Month(), now.Day(), closeTime.Hour(), closeTime.Minute(), closeTime.Second(), 0, time.UTC)

	// 終了時間が開始時間よりも前の場合、終了時間を次の日に設定
	if closeTime.Before(openTime) {
		closeTime = closeTime.Add(24 * time.Hour)
		if now.Before(openTime) {
			now = now.Add(24 * time.Hour)
		}
	}

	return now.After(openTime) && now.Before(closeTime)
}

//TODO: 営業日時のテーブル設計再検討
/*
本日の曜日ごとに、今が営業時間を取得
*/
func resolveBuisnessTime(openingHousr googleplace.OpeningHours, timeD sql.NullTime, isOpen bool) string {

	if timeD.Valid {
		return timeD.Time.Format("15:04:05")
	}
	now := time.Now()
	todayWeekDay := int(now.Weekday())
	fmt.Println("今日の曜日", todayWeekDay)

	var weekPeriodsInfos *googleplace.OpeningHoursPeriodInfo

	for _, v := range openingHousr.Periods {
		var periodInfo googleplace.OpeningHoursPeriodInfo
		if isOpen {
			periodInfo = v.Open
		} else {
			periodInfo = v.Close
		}
		if periodInfo.Day == todayWeekDay {
			weekPeriodsInfos = &periodInfo
			break
		}
	}
	//入ってない場合は、定休日
	if weekPeriodsInfos == nil {
		return "定休日"
	}

	return fmt.Sprintf("%02d:%02d:00", weekPeriodsInfos.Hour, weekPeriodsInfos.Minute)
}

func (h *dogrunHandler) GetDogrunByID(id string) {
	fmt.Println(h.drr.GetDogrunByID(id))
}

/*
指定範囲内のドッグラン検索
*/
func (h *dogrunHandler) SearchAroundDogruns(c echo.Context, condition dto.SearchAroudRectangleCondition) ([]googleplace.BaseResource, error) {
	logger := log.GetLogger(c).Sugar()
	logger.Debugw("検索条件", "condition", condition)

	payload := googleplace.ConvertReqToSearchTextPayload(condition)
	// バリデータのインスタンス作成
	validate := validator.New()
	// カスタムバリデーションルールの登録
	_ = validate.RegisterValidation("latitude", dto.VLatitude)
	_ = validate.RegisterValidation("longitude", dto.VLongitude)
	logger.Infow("リクエストボディ", "payload", payload)

	//base情報のFieldを使用
	var baseFiled googleplace.IFieldMask = googleplace.BaseField{}

	//place情報の取得
	dogrunG, err := h.searchTextUpToSpecifiedTimes(c, payload, baseFiled)
	if err != nil {
		return nil, err
	}
	logger.Infow("レスポンス", "response", dogrunG)
	return dogrunG, nil
}

/*
searchTextを指定上限回数まで実行する
条件：nextPageTokenが含まれている && リクエスト上限回数を超えていないこと
*/
func (h *dogrunHandler) searchTextUpToSpecifiedTimes(c echo.Context, payload googleplace.SearchTextPayLoad, fields googleplace.IFieldMask) ([]googleplace.BaseResource, error) {
	logger := log.GetLogger(c).Sugar()
	var dogruns []googleplace.BaseResource

	var times int
	for {
		times++
		res, err := h.rest.POSTSearchText(c, payload, fields)
		if err != nil {
			return nil, err
		}
		// JSONデータを構造体にデコード
		var searchTextRes googleplace.SearchTextBaseResource
		err = json.Unmarshal(res, &searchTextRes)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		logger.Infow(fmt.Sprintf("レスポンス-%d", times), "response", searchTextRes)
		dogruns = append(dogruns, searchTextRes.Places...)

		//nextPageTokenがない時 or リクエスト上限に達した時
		if searchTextRes.NextPageToken == nil || times > SEARCH_TEXT_MAX_REQUEST_TIMES {
			break
		}

		payload.PageToken = *searchTextRes.NextPageToken
	}

	return dogruns, nil
}
