package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
)

func strToInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return v
}

func parsedSection(str string) []int {
	section := strings.Split(str, "-")
	start := strToInt(section[0])
	end := strToInt(section[1])
	return []int{start, end}
}

func main() {
	day := "04"
	year := "2022"
	useExample := false

	var pt1Ans int
	var pt2Ans int

	inputFile, err := utils.InputFile(day, year, useExample)
	if err != nil {
		panic(err)
	}

	data, err := utils.InputData(inputFile)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(data, "\n")
	for i := range lines {
		pair := strings.Split(lines[i], ",")
		p1Section := parsedSection(pair[0])
		p2Section := parsedSection(pair[1])

		p1SecDiff := p1Section[1] - p1Section[0]
		p2SecDiff := p2Section[1] - p2Section[0]

		// Part 1
		if p1SecDiff >= p2SecDiff &&
			p2Section[0] >= p1Section[0] && p2Section[1] <= p1Section[1] {
			pt1Ans += 1
		} else if p2SecDiff > p1SecDiff &&
			p1Section[0] >= p2Section[0] && p1Section[1] <= p2Section[1] {
			pt1Ans += 1
		}

		// Part 2
		if p2Section[1] >= p1Section[1] && p1Section[1] >= p2Section[0] {
			pt2Ans += 1
		} else if p1Section[1] >= p2Section[1] && p2Section[1] >= p1Section[0] {
			pt2Ans += 1
		}
	}

	fmt.Println("Part 1: ", pt1Ans)
	fmt.Println("Part 2: ", pt2Ans)
}
