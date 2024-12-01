package main

import "testing"

func TestPart1(t *testing.T) {
	input := `
3   4
4   3
2   5
1   3
3   9
3   3
`

    result := Part1(input)
	expected := 11
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `
3   4
4   3
2   5
1   3
3   9
3   3
`

    result := Part2(input)
	expected := 31
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
