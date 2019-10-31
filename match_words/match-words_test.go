package main

import "testing"

func Test_matchWords(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "1",
			args: "__END_DNE-----",
			want: true,
		},
		{
			name: "2",
			args: "__ENDDNE__",
			want: false,
		},
		{
			name: "3",
			args: "IF()()fi[]",
			want: true,
		},
		{
			name: "4",
			args: "for__if__rof__fi",
			want: false,
		},
		{
			name: "5",
			args: "%%$@$while  try ! yrt  for if_fi rof #*#  elihw",
			want: true,
		},
		{
			name: "6",
			args: "",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchWords(tt.args); got != tt.want {
				t.Errorf("matchWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isReverse(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is reverse",
			args: args{
				first: "hello",
				second: "olleh",
			},
			want: true,
		},
		{
			name: "is not reverse",
			args: args{
				first: "code",
				second: "code",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isReverse(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("isReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}