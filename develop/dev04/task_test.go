package dev04

import (
	"reflect"
	"testing"
)

func TestRunSort(t *testing.T) {
	tests := []struct {
		name string
		args *[]string
		want map[string][]string
	}{
		{
			name: "test-0",
			args: &[]string{""},
			want: map[string][]string{},
		},
		{
			name: "test-1",
			args: &[]string{"листок"},
			want: map[string][]string{},
		},
		{
			name: "test-2",
			args: &[]string{"листок", "столик"},
			want: map[string][]string{
				"листок": {"столик"},
			},
		},
		{
			name: "test-3",
			args: &[]string{"листок", "пятак", "пятка", "слиток", "столик", "тяпка"},
			want: map[string][]string{
				"пятак":  {"пятка", "тяпка"},
				"листок": {"слиток", "столик"},
			},
		},
		{
			name: "test-4",
			args: &[]string{"листок", "Пятак", "пятка", "слитоК", "а", "СТолик", "qw", "e", "тяпка"},
			want: map[string][]string{
				"пятак":  {"пятка", "тяпка"},
				"листок": {"слиток", "столик"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunSort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
