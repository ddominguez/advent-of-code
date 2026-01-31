package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var shapes = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var shapePoints = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var outcomePoints = map[string]int{
	"loss": 0,
	"draw": 3,
	"win":  6,
}

var myResult = map[string]string{
	"X": "loss",
	"Y": "draw",
	"Z": "win",
}

// opponents shape mapped to my winning shape
var winningShape = map[string]string{
	"rock":     "paper",
	"paper":    "scissors",
	"scissors": "rock",
}

// opponents shape mapped to my losing shape
var losingShape = map[string]string{
	"rock":     "scissors",
	"paper":    "rock",
	"scissors": "paper",
}

// gameResult will return my result in the game
func gameResult(opp, me string) string {
	if opp == me {
		return "draw"
	}

	if losingShape[opp] == me {
		return "loss"
	}

	return "win"
}

// myShape will return my shape based on the the opponents shape and my expected result
func myShape(opp, res string) string {
	switch res {
	case "loss":
		return losingShape[opp]
	case "win":
		return winningShape[opp]
	default:
		return opp
	}
}

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/02.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	var res int
	for _, round := range strings.Split(strings.TrimSpace(input), "\n") {
		r := strings.Fields(round)
		opp, me := r[0], r[1]
		gameRes := gameResult(shapes[opp], shapes[me])
		res += outcomePoints[gameRes] + shapePoints[shapes[me]]
	}
	return res
}

func part2(input string) int {
	var res int
	for _, round := range strings.Split(strings.TrimSpace(input), "\n") {
		r := strings.Fields(round)
		opp, me := r[0], r[1]
		shape := myShape(shapes[opp], myResult[me])
		res += outcomePoints[myResult[me]] + shapePoints[shape]
	}
	return res
}
