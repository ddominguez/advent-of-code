package main

import "testing"

func TestPart1(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	result := part1(input)
	expected := "CMZ"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

	result := part2(input)
	expected := "MCD"
	if result != expected {
		t.Fatalf("expected %s, got %s", expected, result)
	}
}
