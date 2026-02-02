package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parsedSection(str string) []int {
	section := strings.Split(str, "-")
	start, _ := strconv.Atoi(section[0])
	end, _ := strconv.Atoi(section[1])
	return []int{start, end}
}

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/04.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func process(input string, predicate func(s1, s2 []int) bool) int {
	var res int
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		pair := strings.Split(line, ",")
		section1 := parsedSection(pair[0])
		section2 := parsedSection(pair[1])
		if predicate(section1, section2) {
			res += 1
		}
	}
	return res
}

func rangeFullyContains(s1, s2 []int) bool {
	return (s2[0] >= s1[0] && s2[1] <= s1[1]) || (s1[0] >= s2[0] && s1[1] <= s2[1])
}

func rangesOverlap(s1, s2 []int) bool {
	return s2[1] >= s1[0] && s1[1] >= s2[0]
}

func part1(input string) int {
	return process(input, rangeFullyContains)
}

func part2(input string) int {
	return process(input, rangesOverlap)
}
