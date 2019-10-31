package apple_stocks

import "testing"

func Test_bestProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "descending",
			args: args{
				prices: []int{5, 4, 3, 2, 1},
			},
			want: 0,
		},
		{
			name: "descending",
			args: args{
				prices: []int{0, 20, 5, 2, 21},
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestProfit(tt.args.prices); got != tt.want {
				t.Errorf("bestProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
