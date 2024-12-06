package main

import "testing"

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

func TestPart1(t *testing.T) {
	result := part1(input)
	expected := 41
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(input)
	expected := -100
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}