package main

import "testing"

func Test_knightsJump(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args : args{
				x: 1,
				y: 1,
			},
			want: 2,
		},
		{
			name: "2",
			args : args{
				x: 1,
				y: 2,
			},
			want: 3,
		},
		{
			name: "3",
			args : args{
				x: 1,
				y: 3,
			},
			want: 4,
		},
		{
			name: "3",
			args : args{
				x: 1,
				y: 4,
			},
			want: 4,
		},
		{
			name: "3",
			args : args{
				x: 2,
				y: 4,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightsJump(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("knightsJump() = %v, want %v", got, tt.want)
			}
		})
	}
}