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
	data, _ := os.ReadFile("../../input/07.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	var res int
	cmds := strings.Split(strings.TrimSpace(input), "\n")
	dirSizes := makeDirSizes(cmds)

	for k := range dirSizes {
		if dirSizes[k] <= 100000 {
			res += dirSizes[k]
		}
	}

	return res
}

func part2(input string) int {
	var res int
	cmds := strings.Split(strings.TrimSpace(input), "\n")
	dirSizes := makeDirSizes(cmds)
	unusedLimit := 30000000
	unusedSpace := 70000000 - dirSizes["/"]

	for k := range dirSizes {
		if unusedSpace+dirSizes[k] >= unusedLimit && (res == 0 || res > dirSizes[k]) {
			res = dirSizes[k]
		}
	}

	return res
}

func makeDirSizes(cmds []string) map[string]int {
	var path []string
	m := make(map[string]int)
	for i := range cmds {
		fields := strings.Fields(cmds[i])
		if fields[1] == "cd" {
			if fields[2] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, fields[2])
			}
		} else if fields[1] == "ls" || fields[0] == "dir" {
			continue
		} else {
			size, _ := strconv.Atoi(fields[0])
			for i := range path {
				m[strings.Join(path[:i+1], "")] += size
			}
		}
	}
	return m
}
