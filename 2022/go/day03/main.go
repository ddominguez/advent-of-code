package main

import (
	"fmt"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

var priorities = make(map[string]int)

func strToMap(items string) map[string]int {
	var m = make(map[string]int)
	for i := range items {
		m[fmt.Sprintf("%c", items[i])] = 1
	}
	return m
}

func main() {
	day := "03"
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

	lines := strings.Split(data, "\n")

	for i := range letters {
		l := fmt.Sprintf("%c", letters[i])
		priorities[l] = i + 1
		priorities[strings.ToUpper(l)] = i + 27
	}

	var pt1Sum int
	var pt2Sum int

	var groupSacks []string = nil

	for i := range lines {
		mid := len(lines[i]) / 2
		c1 := lines[i][:mid]
		c2 := lines[i][mid:]

		// part 1
		c1Items := strToMap(c1)
		var commonItem string
		for j := range c2 {
			commonItem = fmt.Sprintf("%c", c2[j])
			if _, ok := c1Items[commonItem]; ok {
				break
			}
		}

		pt1Sum += priorities[commonItem]

		// part 2
		groupSacks = append(groupSacks, lines[i])
		if len(groupSacks) == 3 {
			s1items := strToMap(groupSacks[0])
			s2items := make(map[string]int)
			for i := range groupSacks[1] {
				s := fmt.Sprintf("%c", groupSacks[1][i])
				if _, ok := s1items[s]; ok {
					s2items[s] = 1
				}
			}
			var badge string
			for i := range groupSacks[2] {
				badge = fmt.Sprintf("%c", groupSacks[2][i])
				if _, ok := s2items[badge]; ok {
					break
				}
			}
			pt2Sum += priorities[badge]

			// clear it
			groupSacks = nil
		}
	}

	fmt.Println("Part 1: ", pt1Sum)
	fmt.Println("Part 2: ", pt2Sum)
}
