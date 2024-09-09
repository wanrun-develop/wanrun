package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 認証型エラーコード
const (
	UnAuthorizedCode = "1-1-c"
)

type WRError struct {
	Type       Type
	CauseBy    string
	Msg        string
	InnerError error
}

func (me *WRError) Error() string {
	return fmt.Sprintf("wanrun error: code[%s], message[%s]", me.Type, me.Msg)
}

func (e WRError) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		fmt.Fprintf(f, "  code[%s]", e.Type)
		fmt.Fprintf(f, "  message: %s", e.Msg)
		fmt.Fprintf(f, " ->causeBy: %s", e.CauseBy)
	default:
		fmt.Fprintf(f, "%v", string(c))
	}
}

/*
エラー生成
すでにWRErrorの場合は、そのまま返す
*/
func NewWRError(err error, msg string, ErrorType Type) *WRError {
	if me, ok := err.(*WRError); ok {
		return me
	}
	return &WRError{
		Type:       ErrorType,
		CauseBy:    err.Error(),
		Msg:        msg,
		InnerError: err,
	}
}

/*
カスタムエラーハンドラーミドルウェア
*/
func HttpErrorHandler(err error, c echo.Context) {
	code := 500

	var me *WRError
	if wreer, ok := err.(*WRError); ok {
		me = wreer
		code = mappingError(me)
	}

	res := ErrorRes{
		Code:       me.Type.String(),
		Message:    me.Msg,
		StackTrace: me.CauseBy,
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			err := c.JSON(code, res)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}

/*
内部エラーコードどHTTPエラーコードのマッピング
*/
func mappingError(err *WRError) int {
	errorType := err.Type

	var httpCode int
	switch errorType.ErrorType {
	case CLIENT:
		switch errorType.Service {
		case AUTH:
			httpCode = http.StatusUnauthorized //401
		default:
			httpCode = http.StatusBadRequest //400
		}
	case SERVER:
		httpCode = http.StatusInternalServerError //500
	default:
		httpCode = http.StatusInternalServerError //500
	}

	return httpCode
}
