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

var shapePoints = map[string]uint64{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var outcomePoints = map[string]uint64{
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
	day := "02"
	year := "2022"
	useExample := false

	inputFile, err := utils.InputFile(day, year, useExample)
	if err != nil {
		panic(err)
	}

	data, err := utils.InputData(inputFile)
	if err != nil {
		panic(err)
	}

	split := strings.Split(data, "\n")

	var pt1Score uint64
	var pt2Score uint64
	for i := range split {
		gameSh := strings.Split(split[i], " ")
		opp := gameSh[0]
		me := gameSh[1]

		pt1Res := gameResult(shapes[opp], shapes[me])
		pt1Score += outcomePoints[pt1Res] + shapePoints[shapes[me]]

		mySh := myShape(shapes[opp], myResult[me])
		if mySh == "" {
			panic("could not get my game shape")
		}
		pt2Score += outcomePoints[myResult[me]] + shapePoints[mySh]
	}

	fmt.Println("Part 1: ", pt1Score)
	fmt.Println("Part 2: ", pt2Score)
}
