package transpose_in_place

import (
	"math"
)

/*
Given a 2D array, transpose it 1 clockwise turn in place
	a	b	c	d
	e	f	g	h
	i	j	k	l
	m	n	o	p
to
	m	i	e	a
	n	j	f	b
	o	k	g	c
	p	l	h	d
or
	a	b	c
	d	e	f
	g	h	i
to
	g d a
	h e b
	i	f	c
*/

func TransposeInPlace(n [][]string) [][]string {
	l := len(n)
	// how many times to move in?
	for i := 0; i < int(math.Floor(float64(len(n))/2.0)); i++ {
		// how many times we do swap at this circular band
		for j := 0; j < len(n)-1-(i*2); j++ {
			// get temp of top right
			temp := n[i+j][l-1-i]
			n[i+j][l-1-i] = n[i][i+j]

			// get temp of bottom right
			temp2 := n[l-1-i][l-1-i-j]
			n[l-1-i][l-1-i-j] = temp

			temp = n[l-1-i-j][i]
			n[l-1-i-j][i] = temp2

			n[i][i+j] = temp
		}
	}
	return n
}
