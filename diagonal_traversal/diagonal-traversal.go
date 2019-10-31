package diagonal_traversal

/*
	given a [][]string such as (all of the ints are actually strings)
	[ [ 7  11 14 16 ],
		[	4  8  12 15	],
		[	2	 5	9	 13	],
		[	1	 3	6	 10	] ]

	return a string slice like this
	[]string{"1", "23", "456", "78910", "111213", "1415" , "16"}
*/

func DiagonalTraversal(n [][]string) []string {
	var output []string
	for i := len(n) - 1; i >= 0; i-- {
		word := ""
		for j := 0; j < len(n)-i; j++ {
			word += n[i+j][j]
		}
		output = append(output, word)
	}

	// columns
	// 1 2 3
	for i := 1; i < len(n); i++ {
		// rows
		// 3 2 1
		word := ""
		for j := len(n) - i; j > 0; j-- {
			word += n[len(n)-i-j][len(n)-j]
		}
		output = append(output, word)
	}
	return output
}
