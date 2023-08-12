package main

import (
	"fmt"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
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
	if res == "draw" {
		return opp
	}

	if res == "win" {
		return winningShape[opp]
	}

	if res == "loss" {
		return losingShape[opp]
	}

	return ""
}

func main() {
	var pt1Ans int
	var pt2Ans int

	data, err := utils.InputData(utils.Puzzle{
		Day:        "02",
		Year:       "2022",
		UseExample: false,
	})
	if err != nil {
		panic(err)
	}

	split := strings.Split(strings.TrimSpace(data), "\n")

	for i := range split {
		gameSh := strings.Split(split[i], " ")
		opp := gameSh[0]
		me := gameSh[1]

		pt1Res := gameResult(shapes[opp], shapes[me])
		pt1Ans += outcomePoints[pt1Res] + shapePoints[shapes[me]]

		mySh := myShape(shapes[opp], myResult[me])
		if mySh == "" {
			panic("could not get my game shape")
		}
		pt2Ans += outcomePoints[myResult[me]] + shapePoints[mySh]
	}

	fmt.Println("Part 1: ", pt1Ans)
	fmt.Println("Part 2: ", pt2Ans)
}
