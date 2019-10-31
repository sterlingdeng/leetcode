package main

import "testing"

func Test_coinSum(t *testing.T) {
	type args struct {
		coins  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{coins: []int{1, 2, 5}, target: 5}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinSum(tt.args.coins, tt.args.target); got != tt.want {
				t.Errorf("coinSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
