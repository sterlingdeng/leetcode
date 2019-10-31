package main

/*
nothing fancy, just reverse an input string
try different methods in because go is confusing with runes, chars, etc
*/

func reverse(s string) string {
	var temp rune
	runeslice := []rune(s)
	front := 0
	back := len(runeslice) - 1
	for front < back {
		temp = runeslice[front]
		runeslice[front] = runeslice[back]
		runeslice[back] = temp
		front++
		back--
	}
	return string(runeslice)
}

