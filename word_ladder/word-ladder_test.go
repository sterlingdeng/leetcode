package word_ladder

import (
	"reflect"
	"testing"
)

func Test_makeAdjList(t *testing.T) {
	type args struct {
		wordList []string
	}
	tests := []struct {
		name     string
		args     args
		wordWild map[string][]string
		wildWord map[string][]string
	}{
		{
			name: "test1",
			args: args{
				wordList: []string{
					"hot", "dog", "dot", "lot", "log", "cog",
				},
			},
			wordWild: map[string][]string{
				"hot": {"*ot", "h*t", "ho*"},
				"dog": {"*og", "d*g", "do*"},
				"dot": {"*ot", "d*t", "do*"},
				"lot": {"*ot", "l*t", "lo*"},
				"log": {"*og", "l*g", "lo*"},
				"cog": {"*og", "c*g", "co*"},
			},
			wildWord: map[string][]string{
				"*ot": {"hot", "dot", "lot"},
				"h*t": {"hot"},
				"ho*": {"hot"},
				"*og": {"dog", "log", "cog"},
				"d*g": {"dog"},
				"do*": {"dog", "dot"},
				"d*t": {"dot"},
				"l*t": {"lot"},
				"lo*": {"lot", "log"},
				"l*g": {"log"},
				"c*g": {"cog"},
				"co*": {"cog"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := makeAdjList(tt.args.wordList)
			if !reflect.DeepEqual(got1, tt.wordWild) {
				t.Errorf("wordWild = %v, want %v", got1, tt.wordWild)
			}
			if !reflect.DeepEqual(got2, tt.wildWord) {
				t.Errorf("wildWord = %v, want %v", got2, tt.wildWord)
			}
		})
	}
}

func Test_word_ladder(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				beginWord: "hit",
				endWord: "cog",
				wordList: []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := word_ladder(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("word_ladder() = %v, want %v", got, tt.want)
			}
		})
	}
}