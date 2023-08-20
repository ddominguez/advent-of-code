package main

import (
	"fmt"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
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
	data, err := utils.InputData(utils.Puzzle{
		Day:        "06",
		Year:       "2022",
		UseExample: false,
	})
	if err != nil {
		panic(err)
	}

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
