package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/05.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) string {
	split := strings.Split(input, "\n\n")
	stackMap := makeStackMap(split[0])
	procedures := strings.Split(strings.TrimSpace(split[1]), "\n")

	for i := range procedures {
		fields := strings.Fields(procedures[i])
		move, _ := strconv.Atoi(fields[1])
		from, _ := strconv.Atoi(fields[3])
		to, _ := strconv.Atoi(fields[5])
		moving := stackMap[from][len(stackMap[from])-move : len(stackMap[from])]

		slices.Reverse(moving)
		stackMap[to] = append(stackMap[to], moving...)
		stackMap[from] = stackMap[from][:len(stackMap[from])-move]
	}

	return buildResult(stackMap)
}

func part2(input string) string {
	split := strings.Split(input, "\n\n")
	stackMap := makeStackMap(split[0])
	procedures := strings.Split(strings.TrimSpace(split[1]), "\n")

	for i := range procedures {
		fields := strings.Fields(procedures[i])
		move, _ := strconv.Atoi(fields[1])
		from, _ := strconv.Atoi(fields[3])
		to, _ := strconv.Atoi(fields[5])
		moving := stackMap[from][len(stackMap[from])-move : len(stackMap[from])]
		stackMap[to] = append(stackMap[to], moving...)
		stackMap[from] = stackMap[from][:len(stackMap[from])-move]
	}

	return buildResult(stackMap)
}

func buildResult(m map[int][]string) string {
	var res []string
	for k := 1; k <= len(m); k++ {
		res = append(res, m[k][len(m[k])-1])
	}
	return strings.Join(res, "")
}

func makeStackMap(data string) map[int][]string {
	m := make(map[int][]string)
	rows := strings.Split(data, "\n")
	for i := len(rows) - 2; i >= 0; i-- {
		for j, col := 1, 1; j < len(rows[i]); j += 4 {
			if rows[i][j] >= 'A' && rows[i][j] <= 'Z' {
				m[col] = append(m[col], string(rows[i][j]))
			}
			col += 1
		}
	}
	return m
}
