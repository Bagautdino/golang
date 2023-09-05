package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	sheriff := "sheriff"
	letterCount := make(map[rune]int)
	for _, char := range s {
		char = toLower(char)
		if isSheriffLetter(char) {
			letterCount[char]++
		}
	}
	minCount := -1
	for _, char := range sheriff {
		char = toLower(char)
		charCount := letterCount[char]
		if char == 'f' {
			charCount /= 2
		}
		if minCount == -1 || charCount < minCount {
			minCount = charCount
		}
	}

	fmt.Println(minCount)
}

func toLower(char rune) rune {
	if 'A' <= char && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

func isSheriffLetter(char rune) bool {
	return char == 's' || char == 'h' || char == 'e' || char == 'r' || char == 'i' || char == 'f'
}
