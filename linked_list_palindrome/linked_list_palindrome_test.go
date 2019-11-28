package linked_list_palindrome

import (
	. "hack_hour/data_structures"
	"testing"
)

func Test_isLinkedListPalindrome(t *testing.T) {
	type args struct {
		l *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				l: nil,
			},
			want: true,
		},
		{
			name: "1 element",
			args: args{
				l: &Node{
					Value: 2,
					Next:  nil,
				},
			},
			want: true,
		},
		{
			name: "even and palindrome",
			args: args{
				l: &Node{
					Value: 2,
					Next: &Node{
						Value: 3,
						Next: &Node{
							Value: 3,
							Next: &Node{
								Value: 2,
								Next:  nil,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "odd and palindrome",
			args: args{
				l: &Node{
					Value: 2,
					Next: &Node{
						Value: 3,
						Next: &Node{
							Value: 2,
							Next:  nil,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "odd and not palindrome",
			args: args{
				l: &Node{
					Value: 4,
					Next: &Node{
						Value: 3,
						Next: &Node{
							Value: 2,
							Next:  nil,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "even and not palindrome",
			args: args{
				l: &Node{
					Value: 4,
					Next: &Node{
						Value: 3,
						Next: &Node{
							Value: 2,
							Next:  &Node{
								Value: 4,
								Next:  nil,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLinkedListPalindrome(tt.args.l); got != tt.want {
				t.Errorf("isLinkedListPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
