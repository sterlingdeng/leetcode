package main

import "fmt"

func strstr(s string, x string) int {
	// loops s
	// if the first char in x is found in s
	// iterate through the length of x, if complete return true, if letters dn't match returns false
	for i := 0; i < len(s); i++ {
		// if the first letters match
		if s[i] == x[0] {
			// start at where it matches
			walked := 0
			for j := i; walked < len(x) && j < len(s); j++ {
				fmt.Printf("s: %s, x: %s\n", s[j], x[j-i])
				if rune(s[j]) != rune(x[j-i]) {
					break
				}
				walked++
			}
			if walked == len(x) {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strstr("sst", "st"))
}