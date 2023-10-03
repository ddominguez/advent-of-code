package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
)

type dirItem struct {
	name   string
	size   int
	isDir  bool
	parent string
}

func main() {
	data, err := utils.InputData(utils.Puzzle{
		Day:        "07",
		Year:       "2022",
		UseExample: true,
	})
	if err != nil {
		panic(err)
	}

	var path []string
	dirSizes := make(map[string]int)

	lines := strings.Split(strings.TrimSpace(data), "\n")
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
			size, _ := utils.StrToInt(split[0])
			for i := range path {
				dirSizes[strings.Join(path[:i+1], "")] += size
			}
		}
	}

	part1 := 0
	part2 := 0

	unusedLimit := 30000000
	unusedSpace := 70000000 - dirSizes["/"]
	var sizeValues []int

	for key := range dirSizes {
		if dirSizes[key] <= 100000 {
			part1 += dirSizes[key]
		}
		sizeValues = append(sizeValues, dirSizes[key])
	}

	sort.Slice(sizeValues, func(i, j int) bool { return sizeValues[i] < sizeValues[j] })
	for i := range sizeValues {
		if (unusedSpace + sizeValues[i]) >= unusedLimit {
			part2 = sizeValues[i]
			break
		}
	}

    fmt.Println("Part 1:", part1)
    fmt.Println("Part 2:", part2)
}
