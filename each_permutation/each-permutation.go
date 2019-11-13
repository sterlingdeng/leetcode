package each_permutation

/*
Write a function that takes two arguments, an array and a callback. The callback will be called once for each permutation
of the elements of the array, where the permutation is passed as an argument.
Permutations are all the possible rearrangements of the elements. For this problem, assume each permutation must include
every element from the array.
Permutations must NOT repeat any of the array elements.
No return value is necessary.
eachPermutation([1, 2, 3], function(perm) {
  console.log(perm)
});
[ 1, 2, 3 ]
[ 1, 3, 2 ]
[ 2, 1, 3 ]
[ 2, 3, 1 ]
[ 3, 1, 2 ]
[ 3, 2, 1 ]
*/

func EachPermutation(slice []int) [][]int {
	return helper([]int{}, slice, [][]int{})
}


func helper(currPerm, avail []int, result [][]int) [][]int {
	if len(avail) == 0 {
		return append(result, currPerm)
	}
	for i, val := range avail {
		var newAvail []int
		newAvail = append(newAvail, avail[:i]...)
		newAvail = append(newAvail, avail[i+1:]...)
		newCurrPerm := append(currPerm, val)
		result = helper(newCurrPerm, newAvail, result)
	}
	return result
}

