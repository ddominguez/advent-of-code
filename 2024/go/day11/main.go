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
		fmt.Println(part1(string(data), 75))
	} else {
		fmt.Println(part1(string(data), 25))
	}
}

func part1(input string, count int) int {
	stoneCounts := make(map[string]int)
	for _, s := range strings.Fields(strings.TrimSpace(input)) {
		stoneCounts[s]++
	}
	newStoneCounts, _ := blink(count, stoneCounts)
	var result int
	for _, v := range newStoneCounts {
		result += v
	}
	return result
}

// func part2(input string) int {
// 	var result int
// 	return result
// }

func blink(count int, stones map[string]int) (map[string]int, int) {
	// fmt.Println(stones)
	if count == 0 {
		return stones, count
	}
	updatedStones := make(map[string]int)
	for k, v := range stones {
		for range v {
			for _, s := range changeStones(k) {
				updatedStones[s]++
			}
		}
	}
	return blink(count-1, updatedStones)
}

func changeStones(stone string) []string {
	updated := []string{}
	switch {
	case stone == "0":
		updated = append(updated, "1")
	case len(stone)%2 == 0:
		a := stone[:len(stone)/2]
		b := strings.TrimLeft(stone[len(stone)/2:], "0")
		if b == "" {
			b = "0"
		}
		updated = append(updated, a)
		updated = append(updated, b)
	default:
		n, _ := strconv.Atoi(stone)
		updated = append(updated, strconv.Itoa(n*2024))
	}
	return updated
}
