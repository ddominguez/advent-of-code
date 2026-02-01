package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/03.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func makePriorities() map[string]int {
	p := make(map[string]int)
	for ch := 'a'; ch <= 'z'; ch++ {
		p[string(ch)] = int(ch - 'a' + 1)
		p[string(ch-32)] = int(ch - 'a' + 27)
	}
	return p
}

func strToMap(items string) map[string]struct{} {
	m := make(map[string]struct{})
	for i := range items {
		m[string(items[i])] = struct{}{}
	}
	return m
}

func part1(input string) int {
	var res int
	priorities := makePriorities()

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		mid := len(line) / 2
		c1 := line[:mid]
		c2 := line[mid:]
		c1Items := strToMap(c1)
		for j := range c2 {
			commonItem := string(c2[j])
			if _, ok := c1Items[commonItem]; ok {
				res += priorities[commonItem]
				break
			}
		}
	}

	return res
}

func part2(input string) int {
	var res int
	var groupSacks []string
	priorities := makePriorities()

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		groupSacks = append(groupSacks, line)
		if len(groupSacks) == 3 {
			s1Items := strToMap(groupSacks[0])
			s2Items := make(map[string]struct{})

			for i := range groupSacks[1] {
				s := string(groupSacks[1][i])
				if _, ok := s1Items[s]; ok {
					s2Items[s] = struct{}{}
				}
			}

			for i := range groupSacks[2] {
				badge := string(groupSacks[2][i])
				if _, ok := s2Items[badge]; ok {
					res += priorities[badge]
					break
				}
			}

			groupSacks = nil
		}
	}

	return res
}
