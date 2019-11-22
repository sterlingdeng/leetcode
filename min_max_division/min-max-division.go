package min_max_division

import "math"

// https://app.codility.com/programmers/lessons/14-binary_search_algorithm/min_max_division/
// k is the number of times to divide
// monotonic functions with binary search
func MinMaxDivision(k int, arr []int) int {
	var minSum int
	var maxSum int

	for _, val := range arr {
		maxSum += val
		if val > minSum {
			minSum = val
		}
	}
	possible := minSum
	for minSum <= maxSum {
		mid := int(math.Round(float64(minSum + maxSum)/2.0))
		ok := checkDivisible(mid, k, arr)
		if ok {
			possible = mid
			maxSum = mid - 1
		} else {
			minSum = mid + 1
		}
	}
	return possible
}

func checkDivisible(mid, k int, arr []int) bool {
	numBlockAllowed := k
	currentBlockSum := 0
	for i := 0; i < len(arr); i++ {
		currentBlockSum += arr[i]
		if currentBlockSum > mid {
			numBlockAllowed--
			currentBlockSum = arr[i]
		}
		if numBlockAllowed == 0 {
			return false
		}
	}
	return true
}


