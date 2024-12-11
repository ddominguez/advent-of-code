package main

import "testing"

var input = ``

func TestPart1(t *testing.T) {
    input = "125 17"
    count := 6
	expected := 22
	result := part1(input, count)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

// func TestPart2(t *testing.T) {
// 	result := part2(input)
// 	expected := -1000
// 	if result != expected {
// 		t.Fatalf("expected %d, got %d", expected, result)
// 	}
// }
