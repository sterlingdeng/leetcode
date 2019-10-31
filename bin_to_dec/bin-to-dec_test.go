package main

import "testing"

func Test_binToDec(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: "0",
			},
			want: 0,
		},
		{
			name: "test2",
			args: args{
				s: "11",
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				s: "0101",
			},
			want: 5,
		},
		{
			name: "test1",
			args: args{
				s: "101",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binToDec(tt.args.s); got != tt.want {
				t.Errorf("binToDec() = %v, want %v", got, tt.want)
			}
		})
	}
}