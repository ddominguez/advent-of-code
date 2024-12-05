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
			if grid[row][col] != "X" {
				continue
			}
			res += findXMAS(grid, []int{row, col}, width-1, height-1)
		}
	}

	return res
}

func findXMAS(grid [][]string, currPos []int, maxCol, maxRow int) int {
	var count int
	row, col := currPos[0], currPos[1]
	wl := len("XMAS")
	okRight := col+(wl-1) <= maxCol
	okLeft := col-(wl-1) >= 0
	okUp := row-(wl-1) >= 0
	okDown := row+(wl-1) <= maxRow

	if okRight {
		if grid[row][col+1] == "M" &&
			grid[row][col+2] == "A" &&
			grid[row][col+3] == "S" {
			count++
		}
	}

	if okLeft {
		if grid[row][col-1] == "M" &&
			grid[row][col-2] == "A" &&
			grid[row][col-3] == "S" {
			count++
		}
	}

	if okUp {
		if grid[row-1][col] == "M" &&
			grid[row-2][col] == "A" &&
			grid[row-3][col] == "S" {
			count++
		}
		if okRight &&
			grid[row-1][col+1] == "M" &&
			grid[row-2][col+2] == "A" &&
			grid[row-3][col+3] == "S" {
			count++
		}
		if okLeft &&
			grid[row-1][col-1] == "M" &&
			grid[row-2][col-2] == "A" &&
			grid[row-3][col-3] == "S" {
			count++
		}
	}

	if okDown {
		if grid[row+1][col] == "M" &&
			grid[row+2][col] == "A" &&
			grid[row+3][col] == "S" {
			count++
		}
		if okRight &&
			grid[row+1][col+1] == "M" &&
			grid[row+2][col+2] == "A" &&
			grid[row+3][col+3] == "S" {
			count++
		}
		if okLeft &&
			grid[row+1][col-1] == "M" &&
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
