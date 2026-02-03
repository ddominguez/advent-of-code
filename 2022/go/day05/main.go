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

		for i, j := 0, len(moving)-1; i < j; i, j = i+1, j-1 {
			moving[i], moving[j] = moving[j], moving[i]
		}

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
	rowsCount := len(rows)
	for i := 1; i <= rowsCount-1; i++ {
		items := strings.Split(rows[rowsCount-1-i], " ")
		for j, col := 0, 1; j < len(items); col++ {
			if items[j] == "" {
				j += 4
				continue
			}
			m[col] = append(m[col], string(items[j][1]))
			j += 1
		}
	}

	return m
}
