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
	data, _ := os.ReadFile("../../input/02.txt")

	if *part == 2 {
		fmt.Println(Part2(string(data)))
	} else {
		fmt.Println(Part1(string(data)))
	}
}

func toInts(line string) []int {
	var res []int
	for _, x := range strings.Fields(line) {
		n, _ := strconv.Atoi(x)
		res = append(res, n)
	}
	return res
}

func Part1(input string) int {
	var result int
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		levels := toInts(line)
		if isSafe(levels) {
			result++
		}
	}
	return result
}

func isSafe(levels []int) bool {
	var isIncreasing, isDecreasing bool
	minDiff, maxDiff := -3, 3
	for i := 1; i < len(levels); i++ {
		diff := levels[i-1] - levels[i]
		if diff > 0 {
			isIncreasing = true
		} else if diff < 0 {
			isDecreasing = true
		} else {
			return false
		}
		if isIncreasing && isDecreasing {
			return false
		}
		if diff < minDiff || diff > maxDiff {
			return false
		}
	}
	return true
}

func Part2(input string) int {
	var result int
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		levels := toInts(line)
		if isSafe(levels) {
			result++
		} else {
            for i := range levels {
                newLevels := append([]int{}, levels[:i]...)
                newLevels = append(newLevels, levels[i+1:]...)
				if isSafe(newLevels) {
					result++
					break
				}
			}
		}
	}
	return result
}
