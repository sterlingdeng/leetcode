package diagonal_traversal

import (
	"reflect"
	"testing"
)

func TestDiagonalTraversal(t *testing.T) {
	type args struct {
		n [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "pain and suffering",
			args: args {
				n: [][]string{
					{
						"g", "k", "n", "p",
					},
					{
						"d", "h", "l", "o",
					},
					{
						"b", "e", "i", "m",
					},
					{
						"a", "c", "f", "j",
					},
				},
			},
			want: []string{
				"a", "bc", "def", "ghij", "klm", "no", "p",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiagonalTraversal(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiagonalTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}