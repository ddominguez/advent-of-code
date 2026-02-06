package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/08.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	var visible int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rowsNum := len(lines)
	colsNum := len(lines[0])

	for i := range lines {
		if i == 0 || i == rowsNum-1 {
			visible += colsNum
			continue
		}

		for j := range lines[i] {
			if j == 0 || j == colsNum-1 {
				visible += 1
				continue
			}

			visibleTop := true
			for currRowIndex := i - 1; currRowIndex >= 0; currRowIndex-- {
				if lines[i][j] <= lines[currRowIndex][j] {
					visibleTop = false
					break
				}
			}

			visibleRight := true
			for currColIndex := j + 1; currColIndex <= colsNum-1; currColIndex++ {
				if lines[i][j] <= lines[i][currColIndex] {
					visibleRight = false
					break
				}
			}

			visibleBottom := true
			for currRowIndex := i + 1; currRowIndex <= rowsNum-1; currRowIndex++ {
				if lines[i][j] <= lines[currRowIndex][j] {
					visibleBottom = false
					break
				}
			}

			visibleLeft := true
			for currColIndex := j - 1; currColIndex >= 0; currColIndex-- {
				if lines[i][j] <= lines[i][currColIndex] {
					visibleLeft = false
					break
				}
			}

			if visibleTop || visibleRight || visibleBottom || visibleLeft {
				visible += 1
			}
		}
	}

	return visible
}

func part2(input string) int {
	var scenicScore int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rowsNum := len(lines)
	colsNum := len(lines[0])

	for i := range lines {
		if i == 0 || i == rowsNum-1 {
			continue
		}

		for j := range lines[i] {
			if j == 0 || j == colsNum-1 {
				continue
			}

			seenTop := 0
			for currRowIndex := i - 1; currRowIndex >= 0; currRowIndex-- {
				seenTop += 1
				if lines[i][j] <= lines[currRowIndex][j] {
					break
				}
			}

			seenRight := 0
			for currColIndex := j + 1; currColIndex <= colsNum-1; currColIndex++ {
				seenRight += 1
				if lines[i][j] <= lines[i][currColIndex] {
					break
				}
			}

			seenBottom := 0
			for currRowIndex := i + 1; currRowIndex <= rowsNum-1; currRowIndex++ {
				seenBottom += 1
				if lines[i][j] <= lines[currRowIndex][j] {
					break
				}
			}

			seenLeft := 0
			for currColIndex := j - 1; currColIndex >= 0; currColIndex-- {
				seenLeft += 1
				if lines[i][j] <= lines[i][currColIndex] {
					break
				}
			}

			currScore := seenTop * seenRight * seenBottom * seenLeft
			if currScore > scenicScore {
				scenicScore = currScore
			}

		}
	}

	return scenicScore
}
