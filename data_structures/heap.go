package data_structures

import (
	"errors"
	"math"
)

var ErrHeapEmpty = errors.New("heap empty")

type Heap struct {
	heap       []int
	comparator func(a, b int) bool
}

func NewHeap(comparator func(a, b int) bool) *Heap {
	return &Heap{
		comparator: comparator,
	}
}

func (h *Heap) IsEmpty() bool {
	return len(h.heap) == 0
}

func (h *Heap) Length() int {
	return len(h.heap)
}

func (h *Heap) Peek() (int, error) {
	if len(h.heap) == 0 {
		return -1, ErrHeapEmpty
	}
	return h.heap[0], nil
}

func (h *Heap) Add(n int) {
	h.heap = append(h.heap, n)
	h.siftUp()
}

func (h *Heap) Remove() (int, error) {
	if len(h.heap) == 0 {
		return -1, ErrHeapEmpty
	}
	removed := h.heap[0]
	swap(h.heap, 0, len(h.heap)-1)
	h.heap = h.heap[:len(h.heap)-1]
	h.siftDown(0)
	return removed, nil
}

func (h *Heap) siftUp() {
	currIdx := len(h.heap) - 1
	parentIdx := getParent(currIdx)
	for parentIdx >= 0 && h.comparator(h.heap[currIdx], h.heap[parentIdx]) {
		swap(h.heap, currIdx, parentIdx)
		currIdx = parentIdx
		parentIdx = getParent(currIdx)
	}
}

func (h *Heap) siftDown(idx int) {
	heapSize := len(h.heap)
	element := idx
	leftIdx := getLeftChild(idx)
	rightIdx := getRightChild(idx)
	if leftIdx < heapSize && h.comparator(h.heap[leftIdx], h.heap[element]) {
		element = leftIdx
	}
	if rightIdx < heapSize && h.comparator(h.heap[rightIdx], h.heap[element]) {
		element = rightIdx
	}
	if element == idx {
		return
	}
	swap(h.heap, element, idx)
	h.siftDown(element)
}

func getParent(idx int) int {
	parentIdx := (float64(idx) - 1.0) / 2
	return int(math.Floor(parentIdx))
}

func getLeftChild(idx int) int {
	return idx*2 + 1
}

func getRightChild(idx int) int {
	return idx*2 + 2
}

func swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
