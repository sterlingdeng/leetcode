package main

import "fmt"

/* Write a function called 'transpose' that accepts a two dimensional array and returns a transposed version
 * of that array.
 *
 * Example:
 * const twoDimArray = [ ['fred', 'barney'], [30, 40], [true, false] ]
 * console.log(transpose(twoDimArray)); // -> [['fred', 30, true], ['barney', 40, false]]
 *
 */

/*
	['fred, 'barney']
	[30,			40    ]
	[true,   false  ]

 [fred, 30,   true ]
 [barney, 40, false]


*/

func transpose(in [][]interface{}) [][]interface{} {
	// r ==> c
	// c ==> r
	new := make([][]interface{}, len(in[0]))
	for i := 0; i < len(new); i++ {
		new[i] = make([]interface{}, len(in))
	}
	for c := 0; c < len(in[0]); c++ {
		for r := 0; r < len(in); r++ {
			new[c][r] = in[r][c]
		}
	}
	return new
}

func main() {
	input := [][]interface{}{
		{
			"fred", "barney",
		},
		{
			30, 40,
		},
		{
			true, false,
		},
	}
	fmt.Println(transpose(input))
}
