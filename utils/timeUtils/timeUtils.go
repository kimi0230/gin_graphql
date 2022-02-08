package timeutils

import (
	"time"
)

var LAYOUT = "2006-01-02 15:04:05"
var LAYOUT2 = "20060102150405"

func NowTime(layout string) string {
	if layout == "" {
		layout = LAYOUT
	}
	timelocal, _ := time.LoadLocation("Asia/Taipei")
	time.Local = timelocal
	t := time.Now().Local()

	return t.Format(layout)
}
