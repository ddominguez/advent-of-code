package main

import "testing"

func TestPart1(t *testing.T) {
	input := `A Y
B X
C Z`
	result := part1(input)
	expected := 15
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `A Y
B X
C Z`
	result := part2(input)
	expected := 12
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
