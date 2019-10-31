package main

import "fmt"

// Write a function that returns an array containing the numbers 1 to NUM. Put "fizz" in place of numbers divisble by 3, "buzz" in place of numbers divisble by 5, and "fizzbuzz" in place of numbers divisble by both 3 and 5
// fizzbuzz(16);  -> [ 1,
//                     2,
//                     'fizz',
//                     4,
//                     'buzz',
//                     'fizz',
//                     7,
//                     8,
//                     'fizz',
//                     'buzz',
//                     11,
//                     'fizz',
//                     13,
//                     14,
//                     'fizzbuzz',
//                     16 ]

func fizzbuzz(n int) []interface{} {
	rtn := make([]interface{}, n)
	for i := 1; i < n+1; i++ {
		if i%15 == 0 {
			rtn[i-1] = "fizzbuzz"
		} else if i%3 == 0 {
			rtn[i-1] = "fizz"
		} else if i%5 == 0 {
			rtn[i-1] = "buzz"
		} else {
			rtn[i-1] = i
		}
	}
	return rtn
}

func main() {
	fmt.Println(fizzbuzz(0))
}
