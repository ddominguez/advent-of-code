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
	data, _ := os.ReadFile("../../input/07.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var result int
	var combos = make(map[int][][]string)
	for _, line := range lines {
		testValue, nums := parse(line)
		var opCombos [][]string
		opsLen := len(nums) - 1
		if combos[opsLen] != nil {
			opCombos = combos[opsLen]
		} else {
			opCombos = generateCombos(opsLen, false)
			combos[opsLen] = opCombos
		}

		if hasValidEquation(nums, opCombos, testValue) {
			result += testValue
		}
	}

	return result
}

func part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var result int
	var combos = make(map[int][][]string)
	for _, line := range lines {
		testValue, nums := parse(line)
		var opCombos [][]string
		opsLen := len(nums) - 1
		if combos[opsLen] != nil {
			opCombos = combos[opsLen]
		} else {
			opCombos = generateCombos(opsLen, true)
			combos[opsLen] = opCombos
		}

		if hasValidEquation(nums, opCombos, testValue) {
			result += testValue
		}
	}

	return result
}

func hasValidEquation(nums []int, opCombos [][]string, testVal int) bool {
	for _, ops := range opCombos {
		calcTotal := 0
		for i, num := range nums {
			if i == 0 {
				calcTotal = num
				continue
			}
			calcTotal = doOp(calcTotal, num, ops[i-1])
		}
		if calcTotal == testVal {
			return true
		}
	}
	return false
}

func doOp(a, b int, op string) int {
	switch op {
	case "*":
		return a * b
	case "||":
		num, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
		return num
	default:
		return a + b
	}
}

func parse(line string) (int, []int) {
	split := strings.Split(line, ":")
	val, _ := strconv.Atoi(split[0])
	nums := []int{}
	for _, v := range strings.Fields(split[1]) {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	return val, nums
}

func generateCombos(length int, p2 bool) [][]string {
	if length <= 0 {
		return nil
	}

	var generate func(current []string, depth int)
	combinations := [][]string{}

	generate = func(current []string, depth int) {
		if depth == length {
			combination := make([]string, len(current))
			copy(combination, current)
			combinations = append(combinations, combination)
			return
		}
		generate(append(current, "+"), depth+1)
		generate(append(current, "*"), depth+1)
		if p2 {
			generate(append(current, "||"), depth+1)
		}
	}

	generate([]string{}, 0)
	return combinations
}
