package main

import (
	"strings"
	"testing"
)

const input = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

const inputb = `
########
#......#
...^...#
########
`

const inputc = `
###
#.#
#.#
#^#
`

func TestPart1(t *testing.T) {
	result := part1(input)
	expected := 41
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(input)
	expected := 6
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestIsLoop(t *testing.T) {
	t.Skip()
	lines := strings.Fields(strings.TrimSpace(input))
	start := PositionDirection{pos: Position{row: 8, col: 7}, dir: "down"}
	obstacle := PositionDirection{pos: Position{row: 9, col: 7}, dir: "down"}
	result := isLoop(lines, start, obstacle)
	expected := true
	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}
