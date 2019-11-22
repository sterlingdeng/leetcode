package min_max_division

import "testing"

func TestMinMaxDivision(t *testing.T) {
	type args struct {
		k   int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				k: 3,
				arr: []int{2, 1, 5, 1, 2, 2, 2},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMaxDivision(tt.args.k, tt.args.arr); got != tt.want {
				t.Errorf("MinMaxDivision() = %v, want %v", got, tt.want)
			}
		})
	}
}