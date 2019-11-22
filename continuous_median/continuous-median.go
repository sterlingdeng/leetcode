package continuous_median

import (
	. "hack_hour/data_structures"
)

type ContinuousMedian struct {
	maxHeap *Heap
	minHeap *Heap
}

func NewContinuousMedian() *ContinuousMedian {
	return &ContinuousMedian{
		maxHeap: NewHeap(func(a, b int) bool { return a < b}),
		minHeap: NewHeap(func(a, b int) bool { return a > b}),
	}
}

func (c *ContinuousMedian) Insert(n int) {
	if c.minHeap.Length() == 0 {
		c.minHeap.Add(n)
		return
	}
	if c.minHeap.Length() <= c.maxHeap.Length() {
		c.maxHeap.Add(n)
		removed, _ := c.maxHeap.Remove()
		c.minHeap.Add(removed)
	} else {
		c.minHeap.Add(n)
		removed, _ := c.minHeap.Remove()
		c.maxHeap.Add(removed)
	}
}

func (c *ContinuousMedian) GetMedian() (float64, error) {
	bigger, _ := c.minHeap.Peek()
	smaller, _ := c.maxHeap.Peek()
	if c.minHeap.Length() == c.maxHeap.Length() {
		return (float64(bigger) + float64(smaller)) / 2.0, nil
	} else if c.minHeap.Length() > c.maxHeap.Length() {
		return float64(bigger), nil
	} else {
		return float64(smaller), nil
	}
}
