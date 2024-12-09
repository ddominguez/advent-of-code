package main

import "testing"

const input = `2333133121414131402`

func TestPart1(t *testing.T) {
	result := part1(input)
	expected := 1928
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(input)
	expected := -1000
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

// func TestPoop(t *testing.T) {
// 	t.Skip()
// 	input := "234567874874903342482349"
// 	result := part1(input)
// 	expected := -1000
// 	if result != expected {
// 		t.Fatalf("expected %d, got %d", expected, result)
// 	}
// }
