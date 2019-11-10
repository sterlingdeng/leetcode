package levenshtein_distance

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: "",
				b: "",
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				a: "leven",
				b: "asen",
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				a: "aven",
				b: "asen",
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				a: "",
				b: "asen",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevenshteinDistance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LevenshteinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}