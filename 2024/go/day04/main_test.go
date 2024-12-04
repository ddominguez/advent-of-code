package main

import "testing"

const input = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func TestPart1(t *testing.T) {
	result := part1(input)
	expected := 18
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(input)
	expected := 9
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
