package main

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{coins: []int{1, 2, 5}, target: 11},
			want: 3,
		},
		{
			name: "test2",
			args: args{coins: []int{2}, target: 3},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.target); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
