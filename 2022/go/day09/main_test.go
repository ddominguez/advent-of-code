package main

import "testing"

func TestPart1(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	result := part1(input)
	expected := 13
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	result := part2(input)
	expected := 36
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
