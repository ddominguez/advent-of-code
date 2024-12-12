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
	data, _ := os.ReadFile("../../input/11.txt")

	if *part == 2 {
		fmt.Println(run(string(data), 75))
	} else {
		fmt.Println(run(string(data), 25))
	}
}

func run(input string, count int) int {
	stones := make(map[int]int)
	for _, s := range strings.Fields(strings.TrimSpace(input)) {
		n, _ := strconv.Atoi(s)
		stones[n]++
	}

	for range count {
		stones = blink(stones)
	}

	var result int
	for i := range stones {
		result += stones[i]
	}
	return result
}

func blink(stones map[int]int) map[int]int {
	updated := make(map[int]int)
	for stone, total := range stones {
		if stone == 0 {
			updated[1] += total
		} else if str := strconv.Itoa(stone); len(str)%2 == 0 {
			n1, _ := strconv.Atoi(str[:len(str)/2])
			n2, _ := strconv.Atoi(str[len(str)/2:])
			updated[n1] += total
			updated[n2] += total
		} else {
			updated[stone*2024] += total
		}
	}
	return updated
}
