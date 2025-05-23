package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dirItem struct {
	name   string
	size   int
	isDir  bool
	parent string
}

func main() {
	data, _ := os.ReadFile("../../input/07.txt")

	var path []string
	dirSizes := make(map[string]int)

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i := range lines {
		split := strings.Split(lines[i], " ")
		if split[1] == "cd" {
			if split[2] == ".." {
				path = path[:len(path)-1]
			} else {
				path = append(path, split[2])
			}
		} else if split[1] == "ls" || split[0] == "dir" {
			continue
		} else {
			size, _ := strconv.Atoi(split[0])
			for i := range path {
				dirSizes[strings.Join(path[:i+1], "")] += size
			}
		}
	}

	part1 := 0
	part2 := 0

	unusedLimit := 30000000
	unusedSpace := 70000000 - dirSizes["/"]

	for key := range dirSizes {
		if dirSizes[key] <= 100000 {
			part1 += dirSizes[key]
		}

		if unusedSpace+dirSizes[key] >= unusedLimit && (part2 == 0 || part2 > dirSizes[key]) {
			part2 = dirSizes[key]
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
