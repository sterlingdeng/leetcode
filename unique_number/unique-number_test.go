package unique_number

import "testing"

func Test_uniqueNumber(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				slice: []int{1, 2, 3, 1, 2, 3, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueNumber(tt.args.slice); got != tt.want {
				t.Errorf("uniqueNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
