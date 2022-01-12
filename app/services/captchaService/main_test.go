package captchaservice

import (
	"testing"
)

var tests = []struct {
	arg1 string
	want CaptchaResponse
}{
	{
		"123",
		CaptchaResponse{},
	},
}

func TestNew(t *testing.T) {
	for _, tt := range tests {
		if got := New(); got != tt.want {
			t.Errorf("got = %v \n want = %v \n", got, tt.want)
		}
	}
}
