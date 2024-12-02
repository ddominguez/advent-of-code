package main

import "testing"

const input = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func TestPart1(t *testing.T) {
	result := Part1(input)
	expected := 2
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := Part2(input)
	expected := 4
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
