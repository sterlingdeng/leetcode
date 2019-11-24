package treasure_island

import (
	"math"
	"testing"
)

func TestTreasureIsland(t *testing.T) {
	type args struct {
		board [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				board: [][]string{
					{"O", "O", "O", "O"},
					{"D", "O", "D", "O"},
					{"O", "O", "O", "O"},
					{"X", "D", "D", "D"},
				},
			},
			want: 5,
		},
		{
			name: "test2",
			args: args{
				board: [][]string{
					{"O", "O", "O", "X"},
				},
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				board: [][]string{
					{"0", "D", "X"},
				},
			},
			want: math.MaxInt64,
		},
		{
			name: "test4",
			args: args{
				board: [][]string{
					{"O", "O", "O", "O"},
					{"D", "O", "D", "O"},
					{"O", "O", "O", "O"},
					{"D", "D", "O", "X"},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TreasureIsland(tt.args.board); got != tt.want {
				t.Errorf("TreasureIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
