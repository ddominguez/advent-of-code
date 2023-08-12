package main

import (
	"fmt"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
)

func parsedSection(str string) []int {
	section := strings.Split(str, "-")
	start, err := utils.StrToInt(section[0])
	if err != nil {
		panic(err)
	}
	end, err := utils.StrToInt(section[1])
	if err != nil {
		panic(err)
	}
	return []int{start, end}
}

func main() {
	var pt1Ans int
	var pt2Ans int

	data, err := utils.InputData(utils.Puzzle{
		Day:        "04",
		Year:       "2022",
		UseExample: false,
	})
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(data), "\n")
	for i := range lines {
		pair := strings.Split(lines[i], ",")
		p1Section := parsedSection(pair[0])
		p2Section := parsedSection(pair[1])

		// Part 1
		if (p2Section[0] >= p1Section[0] && p2Section[1] <= p1Section[1]) ||
			(p1Section[0] >= p2Section[0] && p1Section[1] <= p2Section[1]) {
			pt1Ans += 1
		}

		// Part 2
		if (p2Section[1] >= p1Section[1] && p1Section[1] >= p2Section[0]) ||
			(p1Section[1] >= p2Section[1] && p2Section[1] >= p1Section[0]) {
			pt2Ans += 1
		}
	}

	fmt.Println("Part 1: ", pt1Ans)
	fmt.Println("Part 2: ", pt2Ans)
}
