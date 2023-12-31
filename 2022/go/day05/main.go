package main

import (
	"fmt"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
)

func main() {
	var result string

	data, err := utils.InputData(utils.Puzzle{
		Day:        "05",
		Year:       "2022",
		UseExample: false,
	})
	if err != nil {
		panic(err)
	}
	var stackMap = make(map[int][]string)

	split := strings.Split(data, "\n\n")
	stack := split[0]
	procedures := split[1]

	// build stackMap
	stackRows := strings.Split(stack, "\n")
	stackRowsCount := len(stackRows)
	for i := 1; i <= stackRowsCount-1; i++ {
		items := strings.Split(stackRows[stackRowsCount-1-i], " ")
		for j, col := 0, 1; j < len(items); col++ {
			if items[j] == "" {
				j += 4
				continue
			}
			stackMap[col] = append(stackMap[col], strings.Trim(strings.Trim(items[j], "["), "]"))
			j += 1
		}
	}

	for _, p := range strings.Split(strings.TrimSpace(procedures), "\n") {
		splitProc := strings.Split(p, " ")
		move, err := utils.StrToInt(splitProc[1])
		if err != nil {
			panic(err)
		}
		from, err := utils.StrToInt(splitProc[3])
		if err != nil {
			panic(err)
		}
		to, err := utils.StrToInt(splitProc[5])
		if err != nil {
			panic(err)
		}

		moving := stackMap[from][len(stackMap[from])-move : len(stackMap[from])]

		// reverse moving for part 1 answer. comment for part 2 answer.
		for i, j := 0, len(moving)-1; i < j; i, j = i+1, j-1 {
			moving[i], moving[j] = moving[j], moving[i]
		}

		stackMap[to] = append(stackMap[to], moving...)
		stackMap[from] = stackMap[from][:len(stackMap[from])-move]
	}

	for k := 1; k <= len(stackMap); k++ {
		result += stackMap[k][len(stackMap[k])-1]
	}

	fmt.Println("result: ", result)
}
