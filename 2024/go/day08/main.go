package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row int
	col int
}

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
	unique := make(map[Position]bool)
	grid, frequencyPositions := parse(input)
	maxCol := len(grid[0]) - 1
	maxRow := len(grid) - 1
	for _, positions := range frequencyPositions {
		for i, currPos := range positions {
			for _, otherPos := range positions[i+1:] {
				diffRow := currPos.row - otherPos.row
				diffCol := currPos.col - otherPos.col
				abovePos := positionAbove(currPos, diffRow, diffCol)
				if isInBounds(abovePos, maxRow, maxCol) {
					unique[abovePos] = true
				}
				belowPos := positionBelow(otherPos, diffRow, diffCol)
				if isInBounds(belowPos, maxRow, maxCol) {
					unique[belowPos] = true
				}
			}
		}
	}
	return len(unique)
}

func part2(input string) int {
	unique := make(map[Position]bool)
	grid, frequencyPositions := parse(input)
	maxCol := len(grid[0]) - 1
	maxRow := len(grid) - 1
	for _, positions := range frequencyPositions {
		for i, currPos := range positions {
			for _, otherPos := range positions[i+1:] {
				diffRow := currPos.row - otherPos.row
				diffCol := currPos.col - otherPos.col
				aboveCollection := collectAbove(currPos, diffRow, diffCol, maxRow, maxCol)
				for _, ap := range aboveCollection {
					unique[ap] = true
				}
				belowCollection := collectBelow(currPos, diffRow, diffCol, maxRow, maxCol)
				for _, bp := range belowCollection {
					unique[bp] = true
				}
			}
		}
	}
	return len(unique)
}

func collectAbove(pos Position, diffRow, diffCol, maxRow, maxCol int) []Position {
	collection := []Position{}
	for isInBounds(pos, maxRow, maxCol) {
		collection = append(collection, pos)
		pos = positionAbove(pos, diffRow, diffCol)
	}
	return collection
}

func collectBelow(pos Position, diffRow, diffCol, maxRow, maxCol int) []Position {
	collection := []Position{}
	for isInBounds(pos, maxRow, maxCol) {
		collection = append(collection, pos)
		pos = positionBelow(pos, diffRow, diffCol)
	}
	return collection
}

func isInBounds(pos Position, maxRow, maxCol int) bool {
	return pos.row >= 0 && pos.row <= maxRow && pos.col >= 0 && pos.col <= maxCol
}

func positionAbove(pos Position, diffRow, diffCol int) Position {
	return Position{row: pos.row - abs(diffRow), col: pos.col + diffCol}
}
func positionBelow(pos Position, diffRow, diffCol int) Position {
	return Position{row: pos.row + abs(diffRow), col: pos.col - diffCol}
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func parse(input string) ([][]string, map[string][]Position) {
	grid := [][]string{}
	frequencyPositions := make(map[string][]Position)
	for ri, rv := range strings.Fields(strings.TrimSpace(input)) {
		cols := strings.Split(rv, "")
		grid = append(grid, cols)
		for ci, cv := range cols {
			if cv == "." {
				continue
			}
			if frequencyPositions[cv] == nil {
				frequencyPositions[cv] = []Position{}
			}
			frequencyPositions[cv] = append(
				frequencyPositions[cv], Position{row: ri, col: ci})
		}
	}

	return grid, frequencyPositions
}
