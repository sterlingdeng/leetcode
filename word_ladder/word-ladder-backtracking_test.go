package word_ladder

import (
	"testing"
)

func Test_isThere(t *testing.T) {
	type args struct {
		end  string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is there",
			args: args{
				end: "abc",
				list: []string{"hello", "bye", "abc", "yes"},
			},
			want: true,
		},
		{
			name: "is not there",
			args: args{
				end: "cool",
				list: []string{"hello", "bye", "abc", "yes"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isThere(tt.args.end, tt.args.list); got != tt.want {
				t.Errorf("isThere() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_differsByOne(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "differs by 1",
			args: args{"hello", "helly"},
			want: true,
		},
		{
			name: "same word",
			args: args{"hello", "hello"},
			want: true,
		},
		{
			name: "differs by more than 1",
			args: args{"hello", "hnlko"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differsByAtMostOne(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("differsByAtMostOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordLadder(t *testing.T) {
	type args struct {
		start string
		end   string
		list  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				start: "hit",
				end: "cog",
				list: []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "does not exist",
			args: args{
				start: "hit",
				end: "blah",
				list: []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 0,
		},
		{
			name: "only 1 length",
			args: args{
				start: "hit",
				end: "hot",
				list: []string{"hot"},
			},
			want: 2,
		},
		{
			name: "no path exists",
			args: args{
				start: "hit",
				end: "xyz",
				list: []string{"cot", "kyt", "hot"},
			},
			want: 0,
		},
		{
			name: "same words",
			args: args{
				start: "hot",
				end: "dog",
				list: []string{"hot", "dog", "dot"},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WordLadder_backtracking(tt.args.start, tt.args.end, tt.args.list); got != tt.want {
				t.Errorf("WordLadder() = %v, want %v", got, tt.want)
			}
		})
	}
}