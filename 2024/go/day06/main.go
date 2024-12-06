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
	data, _ := os.ReadFile("../../input/06.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

type Position struct {
	row int
	col int
}

func part1(input string) int {
	lines := strings.Fields(strings.TrimSpace(input))
	height := len(lines)
	width := len(lines[0])
	guardPos := Position{}
	direction := "up"
	directions := map[string]Position{
		"up":    {row: -1, col: 0},
		"down":  {row: 1, col: 0},
		"right": {row: 0, col: 1},
		"left":  {row: 0, col: -1},
	}

gridLoop:
	for row := range height {
		for col := range width {
			if lines[row][col] == '^' {
				guardPos.row = row
				guardPos.col = col
				break gridLoop
			}
		}
	}

	visited := map[Position]bool{
		guardPos: true,
	}
	for isInBounds(guardPos, height, width) {
		nextPos := Position{
			row: guardPos.row + directions[direction].row,
			col: guardPos.col + directions[direction].col,
		}
		if !isInBounds(nextPos, height, width) {
			break
		}
		if lines[nextPos.row][nextPos.col] == '#' {
			direction = newDirection(direction)
			continue
		}
		visited[nextPos] = true
		guardPos = nextPos
	}
	return len(visited)
}

func isInBounds(position Position, height, width int) bool {
	row := position.row
	col := position.col
	return col >= 0 && col < width && row >= 0 && row < height
}

func newDirection(direction string) string {
	switch direction {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	}
	return direction
}

func part2(input string) int {
	fmt.Println(input)
	var result int
	return result
}
