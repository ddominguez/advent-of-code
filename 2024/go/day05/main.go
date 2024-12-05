package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/05.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	rules, updates := getRulesAndUpdates(strings.TrimSpace(input))
	middleVals := []string{}
	for _, pages := range updates {
		isCorrect := true
	pagesLoop:
		for i := 0; i < len(pages); i++ {
			currPage := pages[i]
			for _, afterCurPage := range pages[i+1:] {
				if !rules[currPage][afterCurPage] {
					isCorrect = false
					break pagesLoop
				}
			}
		}
		if isCorrect {
			middleVals = append(middleVals, pages[len(pages)/2])
		}
	}

	var result int
	for _, v := range middleVals {
		num, _ := strconv.Atoi(v)
		result += num
	}
	return result
}

func getRulesAndUpdates(input string) (map[string]map[string]bool, [][]string) {
	rules := make(map[string]map[string]bool)
	updates := [][]string{}
	sections := strings.Split(input, "\n\n")

	for _, r := range strings.Split(sections[0], "\n") {
		order := strings.Split(r, "|")
		if rules[order[0]] == nil {
			rules[order[0]] = map[string]bool{}
		}
		rules[order[0]][order[1]] = true
	}
	for _, v := range strings.Split(sections[1], "\n") {
		updates = append(updates, strings.Split(v, ","))
	}
	return rules, updates
}

func part2(input string) int {
	rules, updates := getRulesAndUpdates(strings.TrimSpace(input))
	middleVals := []string{}
	wasIncorrect := false
	for i := 0; i < len(updates); {
		isCorrect := true
		pages := updates[i]
		for j, currPage := range pages {
			for _, afterCurPage := range pages[j+1:] {
				if !rules[currPage][afterCurPage] {
					isCorrect = false
					wasIncorrect = true
					pages[j], pages[j+1] = pages[j+1], pages[j]
					updates[i] = pages
					break
				}
			}
		}
		if isCorrect {
			if wasIncorrect {
				middleVals = append(middleVals, pages[len(pages)/2])
				wasIncorrect = false
			}
			i++
		}
	}

	var result int
	for _, v := range middleVals {
		num, _ := strconv.Atoi(v)
		result += num
	}
	return result
}
