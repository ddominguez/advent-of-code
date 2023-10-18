package main

import (
	"fmt"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
)

func main() {
	data, err := utils.InputData(utils.Puzzle{
		Day:        "10",
		Year:       "2022",
		UseExample: false,
	})
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(data), "\n")
	ops := []int{1, 2}
	x := 1
	cc := 0
	ss := 0
	ccheck := 20
	climit := 220

	for i := range lines {
		isNoop := lines[i] == "noop"
		split := strings.Split(lines[i], " ")
		v := 0
		if !isNoop {
			v, _ = utils.StrToInt(split[1])
		}
		for j := range ops {
			cc += 1
			if cc == ccheck {
				ss += cc * x
				if cc < climit {
					ccheck += 40
				}
			}
			if j > 0 {
				x += v
			}
			if isNoop {
				break
			}
		}
	}

    fmt.Println("signal strengths:", ss)
}
