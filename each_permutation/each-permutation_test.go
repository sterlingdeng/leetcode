package each_permutation

import (
	"fmt"
	"testing"
)

func Test_eachPermuation(t *testing.T) {

	cb := func(i interface{}) {
		fmt.Println(i)
	}

	type args struct {
		slice []int
		cb    func(interface{})
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "first",
			args: args{
				slice: []int{1, 2, 3},
				cb:    cb,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eachPermutation(tt.args.slice, tt.args.cb)
		})
	}
}
