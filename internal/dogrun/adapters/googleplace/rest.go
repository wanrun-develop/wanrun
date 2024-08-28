package googleplace

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wanrun-develop/wanrun/configs"
	"github.com/wanrun-develop/wanrun/pkg/log"
	"go.uber.org/zap"
)

var apiKey string

func init() {
	apiKey = configs.FetchCondigStr("google.place.api.key")
}

type IRest interface {
	GETPlaceInfo(c echo.Context, placeId string, field IFieldMask) ([]byte, error)
}
type rest struct{}

func NewRest() IRest {
	return &rest{}
}

/*
GET
google place apiの実行
*/
func (r *rest) GETPlaceInfo(c echo.Context, placeId string, field IFieldMask) ([]byte, error) {
	logger := log.GetLogger(c).Sugar()

	url := urlApiPlaceWPlaceId(placeId)

	req, err := fetchGETReq(c, url)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Goog-FieldMask", field.getValue())
	logger.Info("field mask:", field.getValue())

	resp, err := exec(c, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスの処理
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

/*
標準のGET用http.Requestの生成
*/
func fetchGETReq(c echo.Context, url string) (*http.Request, error) {
	logger := log.GetLogger(c)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Sugar().Error(err)
		// TODO logger
		return nil, err
	}
	attachBaseHeader(req)
	logger.Info("http requestの生成:",
		zap.String("method", req.Method), zap.String("url", req.URL.String()), zap.Any("header", req.Header))
	return req, nil
}

/*
共通のHeaderをhttp.Requestに付与
*/
func attachBaseHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", apiKey)
	req.Header.Set("Accept-Language", "ja")
}

/*
HTTPリクエストの実行
*/
func exec(c echo.Context, req *http.Request) (*http.Response, error) {
	logger := log.GetLogger(c)

	client := &http.Client{
		Timeout: 60 * time.Second, // タイムアウトの設定
	}

	resp, err := client.Do(req)

	if err != nil {
		logger.Sugar().Error(err)
		return nil, err
	}

	// レスポンスのステータスコードに応じてエラーハンドリング
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body) // レスポンスボディを読み取る
		bodyString := string(bodyBytes)

		//ハンドリング
		switch resp.StatusCode {
		case http.StatusForbidden:
			logger.Error("google place api RESPONSE is Request forbidden", zap.Int("api_response_status_code", resp.StatusCode), zap.String("api_response_body", bodyString))
		case http.StatusNotFound:
			logger.Error("google place api RESPONSE is Not Found", zap.Int("api_response_status_code", resp.StatusCode), zap.String("api_response_body", bodyString))
		case http.StatusInternalServerError:
			logger.Error("google place api RESPONSE is Server error", zap.Int("api_response_status_code", resp.StatusCode), zap.String("rapi_esponse_body", bodyString))
		default:
			logger.Error("google place api RESPONSE is Unexpected error", zap.Int("api_response_tatus_code", resp.StatusCode), zap.String("api_response_body", bodyString))
		}
		return nil, errors.New("failed to exec google palce api")
	} else {
		// ステータスコードが200 OKの場合
		logger.Info("google place api Request is successful")
		return resp, nil
	}
}
