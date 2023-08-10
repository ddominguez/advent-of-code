package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
)

func sum(sl []string) (int64, error) {
	var sum int64
	for i := range sl {
		converted, err := strconv.ParseInt(sl[i], 10, 64)
		if err != nil {
			return -1, err
		}
		sum += converted
	}
	return sum, nil
}

func max(sl []int64) int64 {
	var max int64
	for i := range sl {
		if sl[i] > max {
			max = sl[i]
		}
	}
	return max
}

func main() {
	data, err := utils.InputData(utils.Puzzle{
		Day:        "01",
		Year:       "2022",
		UseExample: false,
	})
	if err != nil {
		panic(err)
	}

	split := strings.Split(data, "\n\n")

	calories := make([]int64, len(split))

	for i, elfCals := range split {
		sum, err := sum(strings.Split(elfCals, "\n"))
		if err != nil {
			panic(err)
		}
		calories[i] = sum
	}

	fmt.Println("Part 1: ", max(calories))

	// Part 2: Sort calories in descending order and get sum of top 3
	sort.Slice(calories, func(i, j int) bool { return int(calories[i]) > int(calories[j]) })

	var top3sum int64
	for _, v := range calories[0:3] {
		top3sum += v
	}
	fmt.Println("Part 2: ", top3sum)
}
