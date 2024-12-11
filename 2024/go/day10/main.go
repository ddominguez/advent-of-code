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
	var result int
	for _, currPos := range startPositions {
		visited := make(Visited)
		currentPath := Positions{}
		score, _, _ := trailHeadScore(grid, currPos, visited, currentPath, 0)
		result += score
	}
	return result
}

func part2(input string) int {
	var result int
	return result
}

func trailHeadScore(grid [][]string, pos Position,
	visited Visited, path Positions, score int) (int, Visited, Positions) {
	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1
	if pos.row < 0 || pos.row > maxRow || pos.col < 0 || pos.col > maxCol {
		return score, visited, path
	}

	if visited[pos] {
		return score, visited, path
	}

	if grid[pos.row][pos.col] == "9" {
		visited[pos] = true
		path = append(path, pos)
		return score + 1, visited, path
	}

	path = append(path, pos)

	for i := range directions {
		gridVal, _ := strconv.Atoi(grid[pos.row][pos.col])
		newPos := Position{
			row: pos.row + directions[i].row,
			col: pos.col + directions[i].col,
		}
		if newPos.row < 0 || newPos.row > maxRow || newPos.col < 0 || newPos.col > maxCol {
			continue
		}
		if visited[newPos] {
			continue
		}
		newVal, _ := strconv.Atoi(grid[newPos.row][newPos.col])
		if gridVal+1 != newVal {
			continue
		}
		score, visited, path = trailHeadScore(grid, newPos, visited, path, score)
	}

	path, last := pop(path)
	visited[last] = true
	return trailHeadScore(grid, last, visited, path, score)
}

func parse(input string) ([][]string, Positions) {
	var grid [][]string
	var startPositions []Position

	for ri, rv := range strings.Fields(strings.TrimSpace(input)) {
		cols := strings.Split(rv, "")
		grid = append(grid, cols)
		for ci := range cols {
			if cols[ci] == "0" {
				startPositions = append(startPositions,
					Position{row: ri, col: ci})
			}
		}
	}

	return grid, startPositions
}

func pop(s Positions) (Positions, Position) {
	if len(s) == 0 {
		return s, Position{}
	}
	return s[:len(s)-1], s[len(s)-1]
}
