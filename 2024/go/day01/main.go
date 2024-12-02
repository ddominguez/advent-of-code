package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func get_data(lines []string, use_frequency bool) ([]int, []int, map[int]int) {
	var left []int
	var right []int
	rightFrequency := make(map[int]int)
	for _, line := range lines {
		pair := strings.Fields(line)
		a, _ := strconv.Atoi(pair[0])
		b, _ := strconv.Atoi(pair[1])
		left = append(left, a)
		if use_frequency {
			rightFrequency[b] += 1
		} else {
			right = append(right, b)
		}
	}
	return left, right, rightFrequency
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	left, right, _ := get_data(lines, false)
	slices.Sort(left)
	slices.Sort(right)

	var total int
	for i, a := range left {
		total += abs(a - right[i])
	}
	return total
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	left, _, rightFrequency := get_data(lines, true)

	var total int
	for _, val := range left {
		total += val * rightFrequency[val]
	}
	return total
}

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/01.txt")

	if *part == 2 {
		fmt.Println(Part2(string(data)))
	} else {
		fmt.Println(Part1(string(data)))
	}
}
