package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("../../input/08.txt")

	visible := 0
	scenicScore := 0

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	linesLen := len(lines)

	for i := range lines {
		if i == 0 || i == linesLen-1 {
			visible += len(lines[i])
			continue
		}
		rowLen := len(lines[i])
		for j := range lines[i] {
			if j == 0 || j == rowLen-1 {
				visible += 1
				continue
			}

			visibleTop := true
			seenTop := 0
			for currRowIndex := i - 1; currRowIndex >= 0; currRowIndex-- {
				seenTop += 1
				if string(lines[i][j]) <= string(lines[currRowIndex][j]) {
					visibleTop = false
					break
				}
			}

			visibleRight := true
			seenRight := 0
			for currColIndex := j + 1; currColIndex <= rowLen-1; currColIndex++ {
				seenRight += 1
				if string(lines[i][j]) <= string(lines[i][currColIndex]) {
					visibleRight = false
					break
				}
			}

			visibleBottom := true
			seenBottom := 0
			for currRowIndex := i + 1; currRowIndex <= linesLen-1; currRowIndex++ {
				seenBottom += 1
				if string(lines[i][j]) <= string(lines[currRowIndex][j]) {
					visibleBottom = false
					break
				}
			}

			visibleLeft := true
			seenLeft := 0
			for currColIndex := j - 1; currColIndex >= 0; currColIndex-- {
				seenLeft += 1
				if string(lines[i][j]) <= string(lines[i][currColIndex]) {
					visibleLeft = false
					break
				}
			}

			if visibleTop || visibleRight || visibleBottom || visibleLeft {
				visible += 1
			}

			if seenTop*seenRight*seenBottom*seenLeft > scenicScore {
				scenicScore = seenTop * seenRight * seenBottom * seenLeft
			}

		}
	}

	fmt.Println("Part 1:", visible)
	fmt.Println("Part 2:", scenicScore)
}
