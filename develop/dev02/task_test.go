package dev02

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "test0",
			arg:  " ",
			want: " ",
		},
		{
			name: "test1",
			arg:  "a4bc2d5e",
			want: "aaaabccddddde",
		},
		{
			name: "test2",
			arg:  "45",
			want: "Error",
		},
		{
			name: "test3",
			arg:  "abcd",
			want: "abcd",
		},
	}
	for _, mytest := range tests {
		t.Run(mytest.name, func(t *testing.T) {
			if got := Unpack(mytest.arg); got != mytest.want {
				t.Errorf("Unpack(): %v, want %v", got, mytest.want)
			}
		})
	}
}
