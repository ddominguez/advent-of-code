package main

import "testing"

func TestPart1(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	expected := 21
	result := part1(input)
	if result != expected {
		t.Fatalf("expected %d, got %d result", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	expected := 8
	result := part2(input)
	if result != expected {
		t.Fatalf("expected %d, got %d result", expected, result)
	}
}
