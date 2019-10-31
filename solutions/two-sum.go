package hack_hour

/*  Given an array of numbers and a target number, return true if there are two numbers in the
 *  array that sum up to the target value; return false otherwise
 */

// brute force
// practice writing for loops like this, since i never do with go
func TwoSum_brute(slice []int, target int) bool {
  for i := 0; i < len(slice); i++ {
  	for j := i+1; j < len(slice); j++ {
  		if slice[i] + slice[j] == target {
  			return true
			}
		}
	}
	return false
}

func TwoSum(slice []int, target int) bool {
	// put the values of the slice in a set
	// no native sets in go, so use a map
	set := make(map[int]struct{}, len(slice))
	for _, val := range slice {
		set[val] = struct{}{}
	}
	for _, val := range slice {
		if _, exists := set[target - val]; exists {
			return true
		}
	}
	return false
}