package main

import "fmt"

func QuickSort(in []int) []int {
	return quicksort(in, 0, len(in)-1)
}

func quicksort(in []int, low, high int) []int {
	//
	if low < high {
		partitionIdx := partition(in, low, high)
		quicksort(in, low, partitionIdx-1)
		quicksort(in, partitionIdx+1, high)
	}
	return in
}

func partition(in []int, low, high int) int {
	pivotVal := in[high]
	ptr := low
	pId := low
	for ptr < high {
		if in[ptr] < pivotVal {
			// swap
			t := in[ptr]
			in[ptr] = in[pId]
			in[pId] = t
			pId++
		}
		ptr++
	}
	//swap pivot value and the value at partition index
	t := in[pId]
	in[pId] = pivotVal
	in[high] = t
	return pId
}

func main() {
	unsorted := []int{5, 3, 8, 20, 19, 1, 2}
	QuickSort(unsorted)
	fmt.Println(unsorted)
}
