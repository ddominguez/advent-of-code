package main

import (
	"fmt"
	"strings"

	"github.com/ddominguez/advent-of-code/utils"
)

type point struct {
	x int32
	y int32
}

var movingPoints = map[string]point{
	"R": {x: 1, y: 0},
	"L": {x: -1, y: 0},
	"U": {x: 0, y: 1},
	"D": {x: 0, y: -1},
}

var touchingPoints = []point{
	// same
	{x: 0, y: 0},
	// above
	{x: 0, y: 1},
	// below
	{x: 0, y: -1},
	// right
	{x: 1, y: 0},
	// left
	{x: -1, y: 0},
	// right top
	{x: 1, y: 1},
	// right bottom
	{x: 1, y: -1},
	// left top
	{x: -1, y: 1},
	// left bottom
	{x: -1, y: -1},
}

// isTouching will return true if the tail is touching the head
// will return false if not touching
func isTouching(h, t point) bool {
	for _, tp := range touchingPoints {
		if t.x == h.x-tp.x && t.y == h.y-tp.y {
			return true
		}
	}
	return false
}

func visitKey(t point) string {
	return fmt.Sprintf("%d,%d", t.x, t.y)
}

func main() {
	data, err := utils.InputData(utils.Puzzle{
		Day:        "09",
		Year:       "2022",
		UseExample: false,
	})
	if err != nil {
		panic(err)
	}

	tailVisited := make(map[string]int)

	lines := strings.Split(strings.TrimSpace(data), "\n")

	knotsCount := 2 // part 1
	// knotsCount := 10 // part 2
	knotsList := make([]point, knotsCount)
	lastKnotIndex := len(knotsList) - 1
	tailVisited[visitKey(knotsList[lastKnotIndex])] = 1

	for i := range lines {
		split := strings.Split(lines[i], " ")
		direction := split[0]
		steps, _ := utils.StrToInt(split[1])

		for step := 0; step < steps; step++ {
			for k := range knotsList {
				if k == 0 {
					knotsList[k].x += movingPoints[direction].x
					knotsList[k].y += movingPoints[direction].y
					continue
				}
				prev := knotsList[k-1]
				curr := knotsList[k]
				if !isTouching(prev, curr) {
					if prev.x > curr.x {
						knotsList[k].x += 1
					} else if prev.x < curr.x {
						knotsList[k].x -= 1
					}
					if prev.y > curr.y {
						knotsList[k].y += 1
					} else if prev.y < curr.y {
						knotsList[k].y -= 1
					}
					if k == lastKnotIndex {
						tailVisited[visitKey(knotsList[k])] += 1
					}
				}
			}
		}
	}

	fmt.Println("total visited:", len(tailVisited))
}
