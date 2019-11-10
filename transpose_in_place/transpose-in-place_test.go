package transpose_in_place

import (
	"reflect"
	"testing"
)

func TestTransposeInPlace(t *testing.T) {
	type args struct {
		n [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "empty",
			args: args{
				n: [][]string{},
			},
			want: [][]string{},
		},
		{
			name: "1x1",
			args: args{
				n: [][]string{{"a"}},
			},
			want: [][]string{{"a"}},
		},
		{
			name: "2x2",
			args: args{
				n: [][]string{
					{"a", "b"},
					{"c", "d"},
				},
			},
			want: [][]string{
				{"c", "a"},
				{"d", "b"},
			},
		},
		{
			name: "2x2",
			args: args{
				n: [][]string{
					{"a", "b", "c", "d"},
					{"e", "f", "g", "h"},
					{"i", "j", "k", "l"},
					{"m", "n", "o", "p"},
				},
			},
			want: [][]string{
				{"m", "i", "e", "a"},
				{"n", "j", "f", "b"},
				{"o", "k", "g", "c"},
				{"p", "l", "h", "d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransposeInPlace(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransposeInPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
