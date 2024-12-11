package main

import "testing"

// p1 expected=1
var _input = `
0123
1234
8765
9876
`

// p1 expected 36
// p2 expected 81
var input = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func TestPart1(t *testing.T) {
	result := part1(input)
	expected := 36
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(input)
	expected := 81
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
