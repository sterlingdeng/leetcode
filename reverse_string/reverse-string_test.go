package main

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		in string
		want string
	}{
		{"1", "Hello", "olleH"},
		{"2", "He😀", "😀eH"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.in); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}