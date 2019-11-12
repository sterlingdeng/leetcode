package min_number_of_jumps

import "testing"

func TestMinNumberOfJumps(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				array: []int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3},
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				array: []int{3},
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				array: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				array: []int{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinNumberOfJumps(tt.args.array); got != tt.want {
				t.Errorf("MinNumberOfJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
