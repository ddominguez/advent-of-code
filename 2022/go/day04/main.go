package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parsedSection(str string) []int {
	section := strings.Split(str, "-")
	start, _ := strconv.Atoi(section[0])
	end, _ := strconv.Atoi(section[1])
	return []int{start, end}
}

func main() {
	var pt1Ans int
	var pt2Ans int

	data, _ := os.ReadFile("../../input/04.txt")

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
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
