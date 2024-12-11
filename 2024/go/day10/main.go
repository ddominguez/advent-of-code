package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row int
	col int
}

type Visited map[Position]bool
type Positions []Position
type Grid [][]int

var directions = []Position{
	{row: -1, col: 0},
	{row: 1, col: 0},
	{row: 0, col: 1},
	{row: 0, col: -1},
}

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/10.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	grid, startPositions := parse(input)
	getRating := false
	initialScore := 0
	var result int
	for _, currPos := range startPositions {
		visited := make(Visited)
		score, _ := walkTrailHead(grid, currPos, visited, getRating, initialScore)
		result += score
	}
	return result
}

func part2(input string) int {
	grid, startPositions := parse(input)
	getRating := true
	initialRating := 0
	var result int
	for _, currPos := range startPositions {
		visited := make(Visited)
		rating, _ := walkTrailHead(grid, currPos, visited, getRating, initialRating)
		result += rating
	}
	return result
}

func walkTrailHead(grid Grid, pos Position, visited Visited, getRating bool, score int) (int, Visited) {
	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	if visited[pos] {
		return score, visited
	}

	if grid[pos.row][pos.col] == 9 {
		if !getRating {
			visited[pos] = true
		}
		return score + 1, visited
	}

	for i := range directions {
		newPos := Position{
			row: pos.row + directions[i].row,
			col: pos.col + directions[i].col,
		}
		if newPos.row < 0 || newPos.row > maxRow || newPos.col < 0 || newPos.col > maxCol {
			continue
		}
		if grid[newPos.row][newPos.col] == grid[pos.row][pos.col]+1 {
			score, visited = walkTrailHead(grid, newPos, visited, getRating, score)
		}
	}
	return score, visited
}

func parse(input string) (Grid, Positions) {
	var grid Grid
	var startPositions []Position
	for ri, rv := range strings.Fields(strings.TrimSpace(input)) {
		cols := []int{}
		for i := range strings.Split(rv, "") {
			n, _ := strconv.Atoi(string(rv[i]))
			cols = append(cols, n)
		}
		grid = append(grid, cols)
		for ci := range cols {
			if cols[ci] == 0 {
				startPositions = append(startPositions,
					Position{row: ri, col: ci})
			}
		}
	}
	return grid, startPositions
}
