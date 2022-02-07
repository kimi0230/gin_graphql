package errorCode

import "net/http"

/*
code :Http `Error Code`-`分類`-`版本`-`編號`
*/

/* 2xx成功 */
// 200
var SUCCESS = map[string]interface{}{
	"code":     "200-API-V1-0001",
	"status":   true,
	"message":  "正確狀況",
	"httpCode": http.StatusOK,
}

// 202
var ACCEPTED = map[string]interface{}{
	"code":     "202-API-V1-0001",
	"status":   false,
	"message":  "Accepted",
	"httpCode": http.StatusAccepted,
}

/******/

/* 4xx客戶端錯誤 */
// 400
var BAD_REQUEST = map[string]interface{}{
	"code":     "400-API-V1-0001",
	"status":   false,
	"message":  "Bad Request",
	"httpCode": http.StatusBadRequest,
}

var PARAMS_INVALID = map[string]interface{}{
	"code":     "400-API-V1-0002",
	"status":   false,
	"message":  "必填、不可填的參數有誤",
	"httpCode": http.StatusBadRequest,
}

var CAPTCHA_INVALID = map[string]interface{}{
	"code":     "400-API-V1-0003",
	"status":   false,
	"message":  "驗證碼錯誤or過期",
	"httpCode": http.StatusBadRequest,
}

// 401
var UNAUTHORIZED = map[string]interface{}{
	"code":     "401-API-V1-0001",
	"status":   false,
	"message":  "Forbidden Request",
	"httpCode": http.StatusUnauthorized,
}

// 403
var FORBIDDEN = map[string]interface{}{
	"code":     "403-API-V1-0001",
	"status":   false,
	"message":  "Forbidden",
	"httpCode": http.StatusForbidden,
}

// 406
var NOT_ACCEPTABLE = map[string]interface{}{
	"code":     "406-API-V1-0001",
	"status":   false,
	"message":  "Forbidden",
	"httpCode": http.StatusNotAcceptable,
}

var RATE_LIMIT = map[string]interface{}{
	"code":     "406-API-V1-0002",
	"status":   false,
	"message":  "Rate Limit",
	"httpCode": http.StatusNotAcceptable,
}

// 408
var REQUEST_TIMEOUT = map[string]interface{}{
	"code":     "408-API-V1-0001",
	"status":   false,
	"message":  "Timeout",
	"httpCode": http.StatusRequestTimeout,
}

/******/

/* 5xx伺服器錯誤 */
// 500
var INTERAL_SERVER_ERROR = map[string]interface{}{
	"code":     "500-API-V1-0001",
	"status":   false,
	"message":  "DB 寫入失敗",
	"httpCode": http.StatusInternalServerError,
}

var DBWRITEFAIL = map[string]interface{}{
	"code":     "500-API-V1-0002",
	"status":   false,
	"message":  "DB 寫入失敗",
	"httpCode": http.StatusInternalServerError,
}

/******/
