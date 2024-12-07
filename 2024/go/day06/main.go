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

type PositionDirection struct {
	pos Position
	dir string
}

var directions = map[string]Position{
	"up":    {row: -1, col: 0},
	"down":  {row: 1, col: 0},
	"right": {row: 0, col: 1},
	"left":  {row: 0, col: -1},
}

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

func getInitialPosition(lines []string) Position {
	position := Position{}
gridLoop:
	for row := range len(lines) {
		for col := range len(lines[0]) {
			if lines[row][col] == '^' {
				position.row = row
				position.col = col
				break gridLoop
			}
		}
	}
	return position
}

func part1(input string) int {
	lines := strings.Fields(strings.TrimSpace(input))
	initialPosition := getInitialPosition(lines)
	initialDirection := "up"
	visited := getVisitedPositions(lines, PositionDirection{
		pos: initialPosition,
		dir: initialDirection,
	})

	unique := make(map[Position]bool)
	for i := range visited {
		unique[visited[i].pos] = true
	}
	return len(unique)
}

func getVisitedPositions(
	lines []string,
	guard PositionDirection,
) []PositionDirection {
	var visited []PositionDirection
	height := len(lines)
	width := len(lines[0])

	for isInBounds(guard.pos, height, width) {
		visited = append(visited, guard)
		nextPos := Position{
			row: guard.pos.row + directions[guard.dir].row,
			col: guard.pos.col + directions[guard.dir].col,
		}
		if !isInBounds(nextPos, height, width) {
			break
		}
		if lines[nextPos.row][nextPos.col] == '#' {
			guard.dir = newDirection(guard.dir)
			continue
		}
		guard.pos = nextPos
	}
	return visited
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
