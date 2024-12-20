package main

import "testing"

// var input = `
// RRRRIICCFF
// RRRRIICCCF
// VVRRRCCFFF
// VVRCCCJFFF
// VVVVCJJCFE
// VVIVCCJJEE
// VVIIICJJEE
// MIIIIIJJEE
// MIIISIJEEE
// MMMISSJEEE
// `
// var p1 = 1930
// var p2 = 1206

var input = `
AAAA
BBCD
BBCC
EEEC
`
var p1 = 140
var p2 = 80

func TestPart1(t *testing.T) {
	result := part1(input)
	expected := p1
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := part2(input)
	expected := p2
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

