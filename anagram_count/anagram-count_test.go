package anagram_count

import "testing"

func Test_anagram_count(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args {
				s: []string{
					"asb",
					"abs",
					"sba",
					"oiu",
					"nop",
					"opn",
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anagram_count(tt.args.s); got != tt.want {
				t.Errorf("anagram_count() = %v, want %v", got, tt.want)
			}
		})
	}
}