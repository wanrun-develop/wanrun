package googleplace

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/wanrun-develop/wanrun/configs"
)

var apiKey string

func init() {
	apiKey = configs.FetchCondigStr("google.place.api.key")
}

type IRest interface {
	GETPlaceInfo(placeId string, field IFieldMask) ([]byte, error)
}
type Rest struct{}

func NewRest() IRest {
	return &Rest{}
}

/*
GET
google place apiの実行
*/
func (r *Rest) GETPlaceInfo(placeId string, field IFieldMask) ([]byte, error) {
	url := urlApiPlaceWPlaceId(placeId)

	req, err := fetchGETReq(url)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Goog-FieldMask", field.getValue())

	resp, err := exec(req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		fmt.Println("respがnil")
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
func fetchGETReq(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// TODO logger
		return nil, err
	}
	attachBaseHeader(req)
	fmt.Println(req)
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
func exec(req *http.Request) (*http.Response, error) {
	client := &http.Client{
		Timeout: 60 * time.Second, // タイムアウトの設定
	}

	resp, err := client.Do(req)

	if resp.StatusCode != 200 {
		return nil, err
	}

	return resp, nil
}
