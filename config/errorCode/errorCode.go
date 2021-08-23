package errorCode

import "net/http"

var SUCCESS = map[string]interface{}{
	"code":     "2000",
	"message":  "正確狀況",
	"httpCode": http.StatusOK,
}

var BAD_REQUEST = map[string]interface{}{
	"code":     "2400",
	"message":  "BadRequest",
	"httpCode": http.StatusBadRequest,
}

var FORBIDDEN = map[string]interface{}{
	"code":     "4030",
	"message":  "Auth 錯誤",
	"httpCode": http.StatusForbidden,
}

var PARAMS_INVALID = map[string]interface{}{
	"code":     "4030",
	"message":  "必填、不可填的參數有誤",
	"httpCode": http.StatusForbidden,
}
