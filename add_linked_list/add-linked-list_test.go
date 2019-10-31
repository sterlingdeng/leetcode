package hack_hour

import (
	"reflect"
	"testing"
)

func generateLists() []*node {
	listOne := &node{
		value: 2,
		next: &node{
			value: 1,
			next: &node{
				value: 5,
				next:  nil,
			},
		},
	}

	listTwo := &node{
		value: 5,
		next: &node{
			value: 9,
			next: &node{
				value: 2,
				next:  nil,
			},
		},
	}

	listThree := &node{
		value: 5,
		next: &node{
			value: 9,
		},
	}

	return []*node{listOne, listTwo, listThree}
}

func Test_addLinkedList(t *testing.T) {
	type args struct {
		l1 *node
		l2 *node
	}

	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "normal",
			args: args{
				l1: generateLists()[0],
				l2: generateLists()[1],
			},
			want: &node{
				value: 7,
				next: &node{
					value: 0,
					next: &node{
						value: 8,
						next:  nil,
					},
				},
			},
		},
		{
			name: "first_longer",
			args: args{
				l1: generateLists()[0],
				l2: generateLists()[2],
			},
			want: &node{
				value: 7,
				next: &node{
					value: 0,
					next: &node{
						value: 6,
						next:  nil,
					},
				},
			},
		},
		{
			name: "all nil",
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addLinkedList(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
