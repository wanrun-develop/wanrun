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
	"github.com/wanrun-develop/wanrun/pkg/errors"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"github.com/wanrun-develop/wanrun/pkg/util"
)

const (
	SEARCH_TEXT_MAX_REQUEST_TIMES = 3 //searchTextの最大リクエスト数。pageSizeを20に指定すると、20*3=60個まで取得する
)

type IDogrunHandler interface {
	GetDogrunDetail(echo.Context, string) (*dto.DogrunDetailDto, error)
	GetDogrunByID(string)
	SearchAroundDogruns(echo.Context, dto.SearchAroudRectangleCondition) ([]dto.DogrunDetailDto, error)
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

	if dogrunG.ID == "" {
		return nil, errors.NewWRError(nil, "指定されたPlaceIdのデータが存在しません。", errors.NewDogrunClientErrorEType())
	}

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

	if dogrunD.IsEmpty() && dogrunG.IsNotEmpty() {
		return resolveDogrunDetailByOnlyGoogle(dogrunG)
	} else if dogrunD.IsNotEmpty() && dogrunG.IsEmpty() {
		return resolveDogrunDetailByOnlyDB(dogrunD)
	}

	var dogrunManager int
	if dogrunD.DogrunManagerID.Valid {
		dogrunManager = int(dogrunD.DogrunManagerID.Int64)
	}

	return dto.DogrunDetailDto{
		DogrunID:        int(dogrunD.DogrunID.Int64),
		DogrunManagerID: dogrunManager,
		PlaceId:         dogrunG.ID,
		Name:            util.ChooseStringValidValue(dogrunD.Name, dogrunG.DisplayName.Text),
		Address:         resolveDogrunAddress(dogrunG, dogrunD),
		Location: dto.Location{
			Latitude:  dogrunG.Location.Latitude,
			Longitude: dogrunG.Location.Longitude,
		},
		BusinessStatus: dogrunG.BusinessStatus,
		NowOpen:        resolveNowOpening(dogrunG, dogrunD),
		BusinessDay:    int(dogrunD.BusinessDay.Int64), // TODO: 一旦、営業日時のテーブル設計再検討
		Holiday:        int(dogrunD.Holiday.Int64),     // TODO: 一旦、営業日時のテーブル設計再検討
		OpenTime:       resolveBuisnessTime(dogrunG, dogrunD, true),
		CloseTime:      resolveBuisnessTime(dogrunG, dogrunD, false),
		Description:    util.ChooseStringValidValue(dogrunD.Description, ""),
		GoogleRating:   dogrunG.Rating,
		DogrunTags:     resolveDogrunTagInfo(dogrunD), // ドッグランタグ情報
		CreateAt:       &dogrunD.CreateAt.Time,
		UpdateAt:       &dogrunD.UpdateAt.Time,
	}

}

/*
DBにデータがない場合、google情報のみでレスポンスを作成する
*/
func resolveDogrunDetailByOnlyGoogle(dogrunG googleplace.BaseResource) dto.DogrunDetailDto {
	var emptyDogrunD model.Dogrun
	return dto.DogrunDetailDto{
		PlaceId: dogrunG.ID,
		Name:    dogrunG.DisplayName.Text,
		Address: resolveDogrunAddress(dogrunG, emptyDogrunD),
		Location: dto.Location{
			Latitude:  dogrunG.Location.Latitude,
			Longitude: dogrunG.Location.Longitude,
		},
		BusinessStatus: dogrunG.BusinessStatus,
		NowOpen:        resolveNowOpening(dogrunG, emptyDogrunD),
		OpenTime:       resolveBuisnessTime(dogrunG, emptyDogrunD, true),
		CloseTime:      resolveBuisnessTime(dogrunG, emptyDogrunD, false),
		GoogleRating:   dogrunG.Rating,
	}
}

/*
google側にデータがない場合、DB情報のみでレスポンスを作成する
*/
func resolveDogrunDetailByOnlyDB(dogrunD model.Dogrun) dto.DogrunDetailDto {

	var emptyDogrunG googleplace.BaseResource

	return dto.DogrunDetailDto{
		DogrunID:        int(dogrunD.DogrunID.Int64),
		DogrunManagerID: int(dogrunD.DogrunManagerID.Int64),
		PlaceId:         dogrunD.PlaceId.String,
		Name:            dogrunD.Name.String,
		Address:         resolveDogrunAddress(emptyDogrunG, dogrunD),
		Location: dto.Location{
			Latitude:  dogrunD.Latitude.Float64,
			Longitude: dogrunD.Longitude.Float64,
		},
		NowOpen:     resolveNowOpening(emptyDogrunG, dogrunD),
		BusinessDay: int(dogrunD.BusinessDay.Int64), // TODO: 一旦、営業日時のテーブル設計再検討
		Holiday:     int(dogrunD.Holiday.Int64),     // TODO: 一旦、営業日時のテーブル設計再検討
		OpenTime:    resolveBuisnessTime(emptyDogrunG, dogrunD, true),
		CloseTime:   resolveBuisnessTime(emptyDogrunG, dogrunD, false),
		Description: dogrunD.Description.String,
		DogrunTags:  resolveDogrunTagInfo(dogrunD), // ドッグランタグ情報
		CreateAt:    &dogrunD.CreateAt.Time,
		UpdateAt:    &dogrunD.UpdateAt.Time,
	}
}

/*
DBからドッグランタグ情報を取得
*/
func resolveDogrunTagInfo(dogrunD model.Dogrun) []dto.DogrunTagDto {

	var dogrunTagInfos []dto.DogrunTagDto

	if dogrunD.IsDogrunTagEmpty() {
		return dogrunTagInfos
	}

	dogrunTag := dogrunD.DogrunTags

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

	var gPostCodeComponent googleplace.AddressComponent
	for _, v := range dogrunG.AddressComponents {
		if slices.Contains(v.Types, googleplace.ADDRESSCOMPONENT_TYPES_POSTAL_CODE) {
			gPostCodeComponent = v
			break
		}
	}
	var address string
	var postCode string

	if dogrunD.IsNotEmpty() && dogrunG.IsNotEmpty() { //両方にある場合
		address = util.ChooseStringValidValue(dogrunD.Address, dogrunG.ShortFormattedAddress)
		postCode = util.ChooseStringValidValue(dogrunD.PostCode, gPostCodeComponent.LongText)
	} else if dogrunD.IsNotEmpty() { //DBにのみある場合
		address = dogrunD.Address.String
		postCode = dogrunD.PostCode.String
	} else if dogrunG.IsNotEmpty() { //google側にのみ
		address = dogrunG.ShortFormattedAddress
		postCode = gPostCodeComponent.LongText
	}

	return dto.Address{PostCode: postCode, Address: address}
}

//TODO Google側の情報使うか検討中
/*
営業時間から、現在が営業中かを判定
*/
func resolveNowOpening(dogrunG googleplace.BaseResource, dogrunD model.Dogrun) bool {
	if dogrunD.IsEmpty() && dogrunG.IsNotEmpty() { //dがない時
		return dogrunG.OpeningHours.OpenNow
	} else if dogrunD.IsNotEmpty() && dogrunG.IsEmpty() { //gがない時
		//検討中
		return false
	}

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
func resolveBuisnessTime(dogrunG googleplace.BaseResource, dogrunD model.Dogrun, isOpen bool) string {
	var timeD sql.NullTime

	openingHours := dogrunG.OpeningHours

	//DB情報の状態によって、対象の時間を取得
	if dogrunD.IsEmpty() {
		timeD.Valid = false
	} else if isOpen {
		timeD = dogrunD.OpenTime.NullTime
	} else {
		timeD = dogrunD.CloseTime.NullTime
	}

	if timeD.Valid {
		return timeD.Time.Format("15:04:05")
	}

	//regularOpeningHoursが空の場合は不明
	if openingHours.IsEmpty() {
		return "不明"
	}

	now := time.Now()
	todayWeekDay := int(now.Weekday())

	var weekPeriodsInfos *googleplace.OpeningHoursPeriodInfo

	for _, v := range openingHours.Periods {
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
func (h *dogrunHandler) SearchAroundDogruns(c echo.Context, condition dto.SearchAroudRectangleCondition) ([]dto.DogrunDetailDto, error) {
	logger := log.GetLogger(c).Sugar()
	logger.Debugw("検索条件", "condition", condition)

	payload := googleplace.ConvertReqToSearchTextPayload(condition)
	// バリデータのインスタンス作成
	validate := validator.New()
	// カスタムバリデーションルールの登録
	_ = validate.RegisterValidation("latitude", dto.VLatitude)
	_ = validate.RegisterValidation("longitude", dto.VLongitude)

	//base情報のFieldを使用
	var baseFiled googleplace.IFieldMask = googleplace.BaseField{}

	//place情報の取得
	dogrunsG, err := h.searchTextUpToSpecifiedTimes(c, payload, baseFiled)
	if err != nil {
		return nil, err
	}
	logger.Infof("googleレスポンスplace数:%d", len(dogrunsG))

	//DBにある指定場所内のドッグランを取得
	dogrunsD, err := h.drr.GetDogrunByRectanglePointer(c, condition)
	if err != nil {
		return nil, err
	}
	logger.Infof("DBから取得数:%d", len(dogrunsD))

	return trimAroundDogrunDetailInfo(dogrunsG, dogrunsD), nil
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
		var searchTextRes *googleplace.SearchTextBaseResource
		err = json.Unmarshal(res, searchTextRes)
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

/*
検索結果をもとに、レスポンス用のDTOを作成
placeIdで、両方にあるデータと、DBにのみあるデータ等で分ける
*/
func trimAroundDogrunDetailInfo(dogrunsG []googleplace.BaseResource, dogrunsD []model.Dogrun) []dto.DogrunDetailDto {
	//google情報からpalceIdをkeyにmapにまとめる
	dogrunsGWithPalceID := make(map[string]googleplace.BaseResource, len(dogrunsG))
	for _, dogrunG := range dogrunsG {
		dogrunsGWithPalceID[dogrunG.ID] = dogrunG
	}

	//DB情報からpalceIdがあるデータのみ、mapにまとめる
	dogrunsDWithPalceID := make(map[string]model.Dogrun)
	//palceIdがないもののみでまとめる
	var dogrunDWithoutPlaceIdS []model.Dogrun
	for _, dogrunD := range dogrunsD {
		if dogrunD.PlaceId.Valid {
			dogrunsDWithPalceID[dogrunD.PlaceId.String] = dogrunD
		} else {
			dogrunDWithoutPlaceIdS = append(dogrunDWithoutPlaceIdS, dogrunD)
		}
	}

	var dogrunDetailInfos []dto.DogrunDetailDto

	//両方にplaceIdがある情報をDTOにつめる
	for placeId, dogrunGValue := range dogrunsGWithPalceID {
		dogrunDValue, existDogrunD := dogrunsDWithPalceID[placeId]
		if existDogrunD {
			//DBにもある場合、両方からデータの選別
			dogrunDetailInfos = append(dogrunDetailInfos, resolveDogrunDetail(dogrunGValue, dogrunDValue))
			//DogrunsDから削除
			delete(dogrunsDWithPalceID, placeId)
		} else {
			//google側にしかない場合
			dogrunDetailInfos = append(dogrunDetailInfos, resolveDogrunDetailByOnlyGoogle(dogrunGValue))
		}
	}

	//placeIdはあるが、google側にないものをDTOにつめる
	for _, dogrunDValue := range dogrunsDWithPalceID {
		dogrunDetailInfos = append(dogrunDetailInfos, resolveDogrunDetailByOnlyDB(dogrunDValue))
	}

	//placeIdがないDBのみのデータをDTOにつめる
	for _, dogrunDValue := range dogrunDWithoutPlaceIdS {
		dogrunDetailInfos = append(dogrunDetailInfos, resolveDogrunDetailByOnlyDB(dogrunDValue))
	}

	return dogrunDetailInfos
}
