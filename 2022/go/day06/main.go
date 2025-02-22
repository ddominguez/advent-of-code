package main

import (
	"fmt"
	"os"
	"strings"
)

func hasUniqueChars(s string) bool {
	for i, c := range s {
		if i != strings.Index(s, string(c)) {
			return false
		}
	}
	return true
}

func main() {
	input, _ := os.ReadFile("../../input/06.txt")
	data := string(input)

	limit := 4 // Part 1
	// limit := 14 // Part 2
	chars := limit

	for i := 0; i <= len(data); i++ {
		chars = limit + i
		if hasUniqueChars(data[:chars][i:]) {
			break
		}
	}

	fmt.Println("Result: ", chars)
}
