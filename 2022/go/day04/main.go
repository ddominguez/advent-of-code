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

func part1(input string) int {
	var res int
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		pair := strings.Split(line, ",")
		p1Section := parsedSection(pair[0])
		p2Section := parsedSection(pair[1])
		if (p2Section[0] >= p1Section[0] && p2Section[1] <= p1Section[1]) ||
			(p1Section[0] >= p2Section[0] && p1Section[1] <= p2Section[1]) {
			res += 1
		}
	}
	return res
}

func part2(input string) int {
	var res int
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		pair := strings.Split(line, ",")
		p1Section := parsedSection(pair[0])
		p2Section := parsedSection(pair[1])
		if (p2Section[1] >= p1Section[1] && p1Section[1] >= p2Section[0]) ||
			(p1Section[1] >= p2Section[1] && p2Section[1] >= p1Section[0]) {
			res += 1
		}
	}
	return res
}
