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
		name string
		args args
		want map[string][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeAdjList(tt.args.wordList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeAdjList() = %v, want %v", got, tt.want)
			}
		})
	}
}