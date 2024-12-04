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
	data, _ := os.ReadFile("../../input/04.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	width := len(lines[0])
	height := len(lines)
	grid := make([][]string, height)
	for i, row := range lines {
		grid[i] = strings.Split(row, "")
	}

	var res int
	for row := range height {
		for col := range width {
			if grid[row][col] == "X" {
				res += findXMAS(grid, []int{row, col})
			}
		}
	}

	return res
}

func findXMAS(grid [][]string, currPos []int) int {
	var count int
	maxCol := len(grid[0]) - 1
	maxRow := len(grid) - 1
	row, col := currPos[0], currPos[1]
	wl := len("XMAS")

	// check right
	if col+(wl-1) <= maxCol {
		if grid[row][col+1] == "M" &&
			grid[row][col+2] == "A" &&
			grid[row][col+3] == "S" {
			count++
		}
	}

	// check left
	if col-(wl-1) >= 0 {
		if grid[row][col-1] == "M" &&
			grid[row][col-2] == "A" &&
			grid[row][col-3] == "S" {
			count++
		}
	}

	// check up
	if row-(wl-1) >= 0 {
		if grid[row-1][col] == "M" &&
			grid[row-2][col] == "A" &&
			grid[row-3][col] == "S" {
			count++
		}
	}

	// check down
	if row+(wl-1) <= maxRow {
		if grid[row+1][col] == "M" &&
			grid[row+2][col] == "A" &&
			grid[row+3][col] == "S" {
			count++
		}
	}

	// check up right
	if row-(wl-1) >= 0 && col+(wl-1) <= maxCol {
		if grid[row-1][col+1] == "M" &&
			grid[row-2][col+2] == "A" &&
			grid[row-3][col+3] == "S" {
			count++
		}
	}

	// check up left
	if row-(wl-1) >= 0 && col-(wl-1) >= 0 {
		if grid[row-1][col-1] == "M" &&
			grid[row-2][col-2] == "A" &&
			grid[row-3][col-3] == "S" {
			count++
		}
	}

	// check down right
	if row+(wl-1) <= maxRow && col+(wl-1) <= maxCol {
		if grid[row+1][col+1] == "M" &&
			grid[row+2][col+2] == "A" &&
			grid[row+3][col+3] == "S" {
			count++
		}
	}

	// check down left
	if row+(wl-1) <= maxRow && col-(wl-1) >= 0 {
		if grid[row+1][col-1] == "M" &&
            grid[row+2][col-2] == "A" &&
            grid[row+3][col-3] == "S" {
			count++
		}
	}

	return count
}

func part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	width := len(lines[0])
	height := len(lines)
	grid := make([][]string, height)
	for i, row := range lines {
		grid[i] = strings.Split(row, "")
	}

	var res int
	for row := range height {
		for col := range width {
			if grid[row][col] != "A" {
				continue
			}
			if row-1 < 0 || row+1 > height-1 || col-1 < 0 || col+1 > width-1 {
				continue
			}
			a := grid[row-1][col-1] == "M" && grid[row+1][col+1] == "S"
			b := grid[row-1][col-1] == "S" && grid[row+1][col+1] == "M"
			c := grid[row-1][col+1] == "M" && grid[row+1][col-1] == "S"
			d := grid[row-1][col+1] == "S" && grid[row+1][col-1] == "M"
			if (a || b) && (c || d) {
				res++
			}
		}
	}

	return res
}
