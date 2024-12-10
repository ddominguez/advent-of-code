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
	updatedBlocks := moveBlocks(blocks)
	return checksum(updatedBlocks)
}

func part2(input string) int {
	blocks := toBlocks(strings.TrimSpace(input))
	updatedBlocks := moveFiles(blocks)
	return checksum(updatedBlocks)
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

func moveFiles(blocks []string) []string {
	l, r1 := 0, len(blocks)-1
	r2 := r1
	for l < r1 {
		// get file
		if blocks[r1] == "." {
			r1--
			r2 = r1
			continue
		}
		for r1 > l && blocks[r2] == blocks[r1-1] {
			r1--
		}
		fileLen := len(blocks[r1 : r2+1])

		// find first set of empty blocks that can fit current file
		haveEmptyBlocks := false
		for emptyCount := 0; l < r1 && !haveEmptyBlocks; {
			for i := range fileLen {
				if blocks[l+i] == "." {
					emptyCount++
				}
			}
			if emptyCount < fileLen {
				l++
				emptyCount = 0
				continue
			}
			haveEmptyBlocks = true
		}
		if haveEmptyBlocks {
			for i := range fileLen {
				blocks[l+i], blocks[r2-i] = blocks[r2-i], blocks[l+i]
			}
		}

		r1--
		r2 = r1
		l = 0
	}
	return blocks
}

func moveBlocks(blocks []string) []string {
	l, r := 0, len(blocks)-1
	for l < r {
		for blocks[l] != "." && l < r {
			l++
		}
		for blocks[r] == "." && l < r {
			r--
		}
		blocks[l], blocks[r] = blocks[r], blocks[l]
		l++
		r--
	}
	return blocks
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
