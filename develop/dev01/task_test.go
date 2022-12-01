package ntp_time

import (
	"fmt"
	"testing"
)

func TestGetTime(t *testing.T) {
	var tests = []struct {
		name string
	}{
		{name: "test-1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(GetTime())
		})
	}
}
