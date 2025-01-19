package main

import (
	"container/list"
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
			queue := list.New()
			queue.PushBack(Position{col: c, row: r})
			for queue.Len() > 0 {
				front := queue.Front()
				current := queue.Remove(front).(Position)
				if visited[current.row][current.col] {
					continue
				}
				if visited[current.row] == nil {
					visited[current.row] = make(map[int]bool)
				}
				visited[current.row][current.col] = true
				area += 1
				for _, dir := range directions {
					nextRow := current.row + dir.row
					nextCol := current.col + dir.col
					isInBounds := 0 <= nextRow && nextRow < height && 0 <= nextCol && nextCol < width
					if isInBounds && grid[nextRow][nextCol] == grid[current.row][current.col] {
						queue.PushBack(Position{col: nextCol, row: nextRow})
					} else {
						perimeter += 1
					}
				}
			}
			result += perimeter * area
		}
	}
	return result
}

func part2(input string) int {
	var result int
	return result
}

func parse(input string) [][]string {
	var result [][]string
	for _, v := range strings.Fields(strings.TrimSpace(input)) {
		result = append(result, strings.Split(v, ""))
	}
	return result
}
