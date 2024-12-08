package main

import (
	"testing"
)

const input = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func TestPart1(t *testing.T) {
	result := part1(input)
	expected := 3749
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(input)
	expected := 11387
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
    // 10976847399437 too low
}
