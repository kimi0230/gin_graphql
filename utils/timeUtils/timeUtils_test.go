package timeutils

import (
	"fmt"
	"testing"
)

func TestNowTime(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
	}{
		{
			name: "2006-01-02 15:04:05",
			arg1: LAYOUT,
		},
		{
			name: "20060102150405",
			arg1: LAYOUT2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NowTime(tt.arg1)
			fmt.Println(got)
		})
	}
}
