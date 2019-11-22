package data_structures

import (
	"reflect"
	"testing"
)

var MinHeapFn = func(a, b int) bool { return a < b }

var MaxHeapFn = func(a, b int) bool { return a > b }
var MinHeap *Heap

//var MaxHeap Heap

func init() {
	MinHeap = NewHeap(MinHeapFn)
}

func TestHeap_Add(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		heap *Heap
		args args
		want []int
	}{
		{
			name: "add to min heap",
			heap: MinHeap,
			args: args{5},
			want: []int{5},
		},
		{
			name: "add to min heap",
			heap: MinHeap,
			args: args{2},
			want: []int{2, 5},
		},
		{
			name: "add to min heap",
			heap: MinHeap,
			args: args{10},
			want: []int{2, 5, 10},
		},
		{
			name: "add to min heap",
			heap: MinHeap,
			args: args{1},
			want: []int{1, 2, 10, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.heap.Add(tt.args.n)
			if !reflect.DeepEqual(tt.want, tt.heap.heap) {
				t.Errorf("Heap should be %v but got %v", tt.want, tt.heap.heap)
			}
		})
	}
}

func TestHeap_Remove(t *testing.T) {
	maxHeap := NewHeap(MaxHeapFn)
	maxHeap.Add(1)
	if peeked, _ := maxHeap.Peek(); peeked != 1 {
		t.Errorf("Max heap peek should be 1, but got %d", peeked)
	}

	maxHeap.Add(10)
	if peeked, _ := maxHeap.Peek(); peeked != 10 {
		t.Errorf("Max heap peek should be 10, but got %d", peeked)
	}
	maxHeap.Add(4)
	if peeked, _ := maxHeap.Peek(); peeked != 10 {
		t.Errorf("Max heap peek should be 10, but got %d", peeked)
	}

	if removed, _ := maxHeap.Remove(); removed != 10 {
		t.Errorf("Max heap removed should be 10, but got %d", removed)
	}

	if peeked, _ := maxHeap.Peek(); peeked != 4 {
		t.Errorf("Max heap peek should be 4, but got %d", peeked)
	}
	if removed, _ := maxHeap.Remove(); removed != 4 {
		t.Errorf("Max heap removed should be 4, but got %d", removed)
	}

	if peeked, _ := maxHeap.Peek(); peeked != 1 {
		t.Errorf("Max heap peek should be 1, but got %d", peeked)
	}
}
