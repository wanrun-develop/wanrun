package errors

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 認証型エラーコード
const (
	Un_AUTHORIZED_CODE string = "1-1-c"
)

type wrError struct {
	errorContext errorContext
	causeBy      string
	msg          string
	innerError   error
}

func (me *wrError) Error() string {
	return fmt.Sprintf("wanrun error: code[%s], message[%s]", me.errorContext, me.msg)
}

func (e wrError) Format(f fmt.State, c rune) {
	switch c {
	case 'v':
		fmt.Fprintf(f, "  code[%s]", e.errorContext)
		fmt.Fprintf(f, "  message: %s", e.msg)
		fmt.Fprintf(f, " ->causeBy: %s", e.causeBy)
	default:
		fmt.Fprintf(f, "%v", string(c))
	}
}

/*
エラー生成
すでにWRErrorの場合は、そのまま返す
*/
func NewWRError(err error, msg string, errorContext errorContext) *wrError {
	if me, ok := err.(*wrError); ok {
		return me
	}
	return &wrError{
		errorContext: errorContext,
		causeBy:      err.Error(),
		msg:          msg,
		innerError:   err,
	}
}

/*
カスタムエラーハンドラーミドルウェア
*/
func HttpErrorHandler(err error, c echo.Context) {
	code := 500

	var me *wrError
	if wreer, ok := err.(*wrError); ok {
		me = wreer
		code = mappingError(me)
	}

	res := ErrorRes{
		Code:       me.errorContext.String(),
		Message:    me.msg,
		StackTrace: me.causeBy,
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
func mappingError(err *wrError) int {
	errorContext := err.errorContext

	var httpCode int
	switch errorContext.errorType {
	case CLIENT:
		switch errorContext.service {
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
