package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	type args struct {
		slice  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				slice: []int{
					1, 2, 3, 4, 5,
				},
				target: 20,
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				slice: []int{
					-1, -2, -3,
				},
				target: -5,
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				slice: []int{
					-4, 3, 1, 0, 23,
				},
				target: -1,
			},
			want: true,
		}, {
			name: "test4",
			args: args{
				slice: []int{
					2, 3,
				},
				target: 5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.slice, tt.args.target); got != tt.want {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

