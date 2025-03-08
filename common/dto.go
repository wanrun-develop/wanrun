package common

import (
	"fmt"
	"time"
)

type WRTime struct {
	time.Time
}

// フォーマットを yyyy/MM/dd HH:mm:ss に固定
const F_yyyyMMddHHmmss = "2006/01/02 15:04:05"

// WRTime の JSON 出力用メソッドをカスタマイズ
func (ct WRTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", ct.Format(F_yyyyMMddHHmmss))
	return []byte(formatted), nil
}

type PaginationReq struct {
	Count int `json:"count" validate:"min=1,max=60"`
	Page  int `json:"page" validate:"min=1,max=100"`
}
