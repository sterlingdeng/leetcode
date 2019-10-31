package main

import "testing"

func Test_numberOfIslands(t *testing.T) {
	type args struct {
		in [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				in: [][]byte{
					{
						1, 1, 1, 1, 0,
					},
					{
						1, 1, 0, 1, 0,
					},
					{
						1, 1, 0, 0, 0,
					},
					{
						0, 0, 0, 0, 0,
					},
				},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				in: [][]byte{
					{
						1, 1, 0, 0, 0,
					},
					{
						1, 1, 0, 0, 0,
					},
					{
						0, 0, 1, 0, 0,
					},
					{
						0, 0, 0, 1, 1,
					},
				},
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				in: [][]byte{
					{
						1, 1,
					},
					{
						1, 1,
					},
				},
			},
			want: 1,
		},
		{
			name: "test4",
			args: args{
				in: [][]byte{
					{
						0, 0,
					},
					{
						0, 0,
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfIslands(tt.args.in); got != tt.want {
				t.Errorf("numberOfIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}