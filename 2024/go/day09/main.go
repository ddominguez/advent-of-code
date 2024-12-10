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
	bl := blocks
	l, r1 := 0, len(bl)-1
	r2 := r1
	for l < r1 {
		// get file
		if bl[r1] == "." {
			r1--
			r2 = r1
			continue
		}
		for r1 > l && bl[r2] == bl[r1-1] {
			r1--
		}
		file := bl[r1 : r2+1]
		fileLen := len(file)

		// find first empty block that can fit current file
		haveEmptyBlocks := false
		available := strings.Repeat(".", fileLen)
		for l < r1 && strings.Join(bl[l:l+fileLen], "") != available {
			l++
		}
		haveEmptyBlocks = strings.Join(bl[l:l+fileLen], "") == available
		if haveEmptyBlocks {
			for i := range fileLen {
				bl[l+i], bl[r2-i] = bl[r2-i], bl[l+i]
			}
		}

		r1--
		r2 = r1
		l = 0
	}
	return bl
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
