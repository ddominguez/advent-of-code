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
	data, _ := os.ReadFile("../../input/09.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	blocks := toBlocks(strings.TrimSpace(input))
	movedBlocks := moveBlocks(blocks)
	return checksum(movedBlocks)
}

func part2(input string) int {
	var result int
	return result
}

func checksum(blocks []string) int {
	var result int
	for id, block := range blocks {
		if block == "." {
			continue
		}
		n, _ := strconv.Atoi(string(block))
		result += id * n
	}
	return result
}

func moveBlocks(blocks []string) []string {
	bl := blocks
	l, r := 0, len(bl)-1
	for l < r {
		for bl[l] != "." && l < r {
			l++
		}
		for bl[r] == "." && l < r {
			r--
		}
		bl[l], bl[r] = bl[r], bl[l]
		l++
		r--
	}
	return bl
}

func toBlocks(input string) []string {
	var result []string
	id := 0
	empty := "."
	for i, c := range input {
		repeatCount, _ := strconv.Atoi(string(c))
		block := ""
		if i%2 == 0 {
			block = strconv.Itoa(id)
			id++
		} else {
			block = empty
		}
		for range repeatCount {
			result = append(result, block)
		}
	}
	return result
}
