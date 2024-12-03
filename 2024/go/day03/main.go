package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/03.txt")

	if *part == 2 {
		fmt.Println(Part2(string(data)))
	} else {
		fmt.Println(Part1(string(data)))
	}
}

func Part1(input string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	re2 := regexp.MustCompile(`\d+`)
	instructions := re.FindAllString(input, -1)

	var res int
	for _, instr := range instructions {
		args := re2.FindAllString(instr, -1)
		res += toInt(args[0]) * toInt(args[1])
	}
	return res
}

func Part2(input string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	re2 := regexp.MustCompile(`\d+`)
	instructions := re.FindAllString(input, -1)
	do := true

	var res int
	for _, instr := range instructions {
		if instr == "don't()" {
			do = false
			continue
		}
		if instr == "do()" {
			do = true
			continue
		}
		if do {
			args := re2.FindAllString(instr, -1)
			res += toInt(args[0]) * toInt(args[1])
		}
	}
	return res
}

func toInt(val string) int {
	num, _ := strconv.Atoi(val)
	return num
}
