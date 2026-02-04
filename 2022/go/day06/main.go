package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/06.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	distinctChars := 4
	return charsProcessed(input, distinctChars)
}

func part2(input string) int {
	distinctChars := 14
	return charsProcessed(input, distinctChars)
}

func charsProcessed(input string, distinctChars int) int {
	counts := make(map[byte]int)
	for i := range input[0:distinctChars] {
		counts[input[i]] += 1
	}

	res := 0
	for lo, hi := 0, distinctChars; hi < len(input); lo, hi = lo+1, hi+1 {
		if isUnique(counts) {
			res = hi
			break
		}

		counts[input[lo]] -= 1
		if counts[input[lo]] == 0 {
			delete(counts, input[lo])
		}
		counts[input[hi]] += 1
	}

	return res
}

func isUnique(c map[byte]int) bool {
	for _, v := range c {
		if v > 1 {
			return false
		}
	}
	return true
}
