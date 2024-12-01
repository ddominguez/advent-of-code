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

func get_data(lines []string) ([]int, []int) {
	var left []int
	var right []int
	for _, line := range lines {
		pair := strings.Fields(line)
		a, _ := strconv.Atoi(pair[0])
		b, _ := strconv.Atoi(pair[1])
		left = append(left, a)
		right = append(right, b)
	}
	return left, right
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	left, right := get_data(lines)
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
	left, right := get_data(lines)

	counter := make(map[int]int)
	for _, val := range right {
		counter[val] += 1
	}

	var total int
	for _, val := range left {
		total += val * counter[val]
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
