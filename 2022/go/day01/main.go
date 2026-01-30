package main

import (
	"cmp"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/01.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int64 {
	calories := caloriesPerElf(strings.Split(strings.TrimSpace(input), "\n\n"))
	return slices.Max(calories)
}

func caloriesPerElf(inventory []string) []int64 {
	calories := make([]int64, len(inventory))
	for i, elfCalories := range inventory {
		var total int64
		for _, v := range strings.Split(elfCalories, "\n") {
			parsed, _ := strconv.ParseInt(v, 10, 64)
			total += parsed
		}
		calories[i] = total
	}
	return calories
}

func part2(input string) int64 {
	calories := caloriesPerElf(strings.Split(strings.TrimSpace(input), "\n\n"))
	slices.SortFunc(calories, func(a, b int64) int {
		return cmp.Compare(b, a)
	})

	var result int64
	for _, v := range calories[0:3] {
		result += v
	}
	return result
}
