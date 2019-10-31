package main

import "testing"

func Test_countStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStairs(tt.args.n); got != tt.want {
				t.Errorf("countStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
