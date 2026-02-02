package main

import "testing"

func TestPart1(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	result := part1(input)
	expected := 2
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	result := part2(input)
	expected := 4
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
