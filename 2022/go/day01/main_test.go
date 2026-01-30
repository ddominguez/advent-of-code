package main

import "testing"

func TestPart1(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	result := part1(input)
	expected := int64(24000)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	result := part2(input)
	expected := int64(45000)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
