package hack_hour

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

func eachPermuation(slice []int, cb func(interface{})) {
	isActive := make([]bool, len(slice))
	// idx refers to the current idx at the perms slice
	idx := 0
	perms := make([]int, len(slice))
	helper(slice, perms, isActive, idx, cb)
}

func helper(slice, perms []int, table []bool, idx int, cb func(interface{})) {
	if len(slice) == idx {
		cb(perms)
		return
	}

	// i refers to the idx on the slice in which the value should be selected and placed
	// in the perms slice
	for i := 0; i < len(slice); i++ {
		if !table[i] {
			perms[idx] = slice[i]
			table[i] = true
			helper(slice, perms, table, idx + 1, cb)
			table[i] = false
		}
	}
}
