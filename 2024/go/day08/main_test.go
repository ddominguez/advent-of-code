package main

import "testing"

const input = `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`

func TestPart1(t *testing.T) {
	result := part1(input)
	expected := 14
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(input)
	expected := -1000
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}