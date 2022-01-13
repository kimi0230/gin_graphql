package captchacontroller

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arg1 int
	want CaptchaResponse
}{
	{
		4,
		CaptchaResponse{},
	},
}

func TestCaptchaString(t *testing.T) {
	for _, tt := range tests {
		if got := CaptchaString(tt.arg1); reflect.DeepEqual(got, CaptchaResponse{}) {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}
