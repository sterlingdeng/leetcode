package main

// given 4 arrays that may contain both numbers and strings
// return a new array with the numbers and/or strings that appear in all 4 arrays
// duplicates are only counted once;
// for example == given the following input

// var array1 = [1, 4, 6, 7, "ferret", 12, 12, 99, 2000, "dog", "dog", 99, 1000];
// var array2 = [15, 9, 9, "ferret", 9, 26, 12, 12, "dog"];
// var array3 = [23, 12, 12, 77, "ferret", 9, 88, 100, "dog"];
// var array4 = ["ferret", 12, 12, 45, 9, 66, 77, 78, 2000];

// your output would be [ 12, 'ferret']

func commonElements(arrays ...[]interface{}) []interface{} {
	have := make(map[interface{}]bool)
	for _, val := range arrays[0] {
		have[val] = true
	}

	for i := 1; i < len(arrays); i++ {
		compare := make(map[interface{}]bool)
		for _, val := range arrays[i] {
			compare[val] = true
		}

		for k := range have {
			if _, exists := compare[k]; !exists {
				delete(have, k)
			}
		}

	}

	// convert map to slice
	var rtn []interface{}
	for k := range have {
		rtn = append(rtn, k)
	}
	return rtn
}