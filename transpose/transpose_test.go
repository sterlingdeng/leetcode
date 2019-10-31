package main

import (
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {
	type args struct {
		in [][]interface{}
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "test1",
			args: args{
				in: [][]interface{}{
					{
						"fred", "barney",
					},
					{
						30, 40,
					},
					{
						true, false,
					},
				},
			},
			want: [][]interface{}{
				{
					"fred", 30, true,
				},
				{
					"barney", 40, false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
