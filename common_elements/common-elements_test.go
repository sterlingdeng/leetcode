package main

import (
	"reflect"
	"testing"
)

func Test_commonElements(t *testing.T) {
	type args struct {
		arrays [][]interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "test1",
			args: args{
				arrays: [][]interface{}{
					{
						1, 4, 6, 7, "ferret", 12, 12, 99, 2000, "dog", "dog", 99, 1000,
					},
					{
						15, 9, 9, "ferret", 9, 26, 12, 12, "dog",
					},
					{
						23, 12, 12, 77, "ferret", 9, 88, 100, "dog",
					},
					{
						"ferret", 12, 12, 45, 9, 66, 77, 78, 2000,
					},
				},
			},
			want: []interface{}{
				12, "ferret",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonElements(tt.args.arrays...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonElements() = %v, want %v", got, tt.want)
			}
		})
	}
}



