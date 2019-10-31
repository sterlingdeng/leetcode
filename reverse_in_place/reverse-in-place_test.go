package reverse_in_place

import (
	"reflect"
	"testing"
)

func Test_reverseInPlace(t *testing.T) {
	type args struct {
		slice []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
    {
    	name: "first",
    	args: args{
    		slice: []rune{'a','b','c'},
			},
			want: []rune{'c','b','a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseInPlace(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseInPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
