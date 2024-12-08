package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var combos = make(map[int][][]string)
var concatCombos = make(map[int][][]string)

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
	var validVals []int
	for _, line := range lines {
		testValue, nums := parse(line)
		var opCombos [][]string
		opsLen := len(nums) - 1
		if combos[opsLen] != nil {
			opCombos = combos[opsLen]
		} else {
			opCombos = generateCombos(opsLen)
			combos[opsLen] = opCombos
		}

		if hasValidEquation(nums, opCombos, testValue) {
			validVals = append(validVals, testValue)
		}
	}

	var result int
	for _, v := range validVals {
		result += v
	}
	return result
}

func part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var validVals []int
	for _, line := range lines {
		testValue, nums := parse(line)
		var opCombos [][]string
		opsLen := len(nums) - 1
		if combos[opsLen] != nil {
			opCombos = combos[opsLen]
		} else {
			opCombos = generateCombos(opsLen)
			combos[opsLen] = opCombos
		}

		if hasValidEquation(nums, opCombos, testValue) {
			validVals = append(validVals, testValue)
		} else {
            // TODO:
			// generate new combos with concat || operator
			// check if valid with concat op combos
		}
	}

	var result int
	for _, v := range validVals {
		result += v
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
	if op == "*" {
		return a * b
	}
	return a + b
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

func generateCombos(length int) [][]string {
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
	}

	generate([]string{}, 0)
	return combinations
}
