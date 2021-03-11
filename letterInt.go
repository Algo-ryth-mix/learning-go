package main

func intToLetters(v int) string {
	var result string

	for x := v; x-1 >= 0; x-- {
		result = string(rune(int('A')+x%26)) + result
		x /= 26
	}
	return result
}
