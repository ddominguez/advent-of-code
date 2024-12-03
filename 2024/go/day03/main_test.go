package main

import "testing"

func TestPart1(t *testing.T) {
	input := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	result := Part1(input)
	expected := 161
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	result := Part2(input)
	expected := 48
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
