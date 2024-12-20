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

type Visited map[int]map[int]bool

var directions = []Position{
	{row: -1, col: 0},
	{row: 1, col: 0},
	{row: 0, col: 1},
	{row: 0, col: -1},
}

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/12.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	grid := parse(input)
	height := len(grid)
	width := len(grid[0])
	visited := make(Visited)

	var result int
	for r := range height {
		for c := range width {
			if visited[r][c] {
				continue
			}
			var perimeter, area int
			current := Position{row: r, col: c}
			// fmt.Printf("--- %s ---\n", grid[r][c])
			perimeter, area, visited = get_perimeter_and_area(grid, current, visited)
			// fmt.Printf("p: %d, a: %d\n", perimeter, area)
			// fmt.Println(visited)
			result += perimeter * area
		}
	}
	return result
}

func part2(input string) int {
	var result int
	return result
}

func get_perimeter_and_area(grid [][]string, current Position, visited Visited) (int, int, Visited) {
	var perimeter int
	var area int
	maxRow := len(grid) - 1
	maxCol := len(grid[0]) - 1

	_, ok := visited[current.row]
	if !ok {
		visited[current.row] = make(map[int]bool)
	}

	area += 1
	visited[current.row][current.col] = true
	target := grid[current.row][current.col]
	for _, dir := range directions {
		next := Position{row: current.row + dir.row, col: current.col + dir.col}
		isInBounds := next.row >= 0 && next.row <= maxRow && next.col >= 0 && next.col <= maxCol
		if isInBounds && !visited[next.row][next.col] && grid[next.row][next.col] == target {
			p, a, v := get_perimeter_and_area(grid, next, visited)
			perimeter += p
			area += a
			visited = v
		}
		if !isInBounds || grid[next.row][next.col] != target {
			perimeter += 1
		}
	}

	return perimeter, area, visited
}

func parse(input string) [][]string {
	var result [][]string
	for _, v := range strings.Fields(strings.TrimSpace(input)) {
		result = append(result, strings.Split(v, ""))
	}
	return result
}
