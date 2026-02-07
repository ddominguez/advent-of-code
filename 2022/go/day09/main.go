package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func abs(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}

func isTouching(h, t point) bool {
	x := abs(h.x - t.x)
	y := abs(h.y - t.y)
	return x <= 1 && y <= 1
}

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/09.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

func part1(input string) int {
	motions := strings.Split(strings.TrimSpace(input), "\n")
	knots := 2
	return positionsVisited(motions, knots)
}

func part2(input string) int {
	motions := strings.Split(strings.TrimSpace(input), "\n")
	knots := 10
	return positionsVisited(motions, knots)
}

func positionsVisited(motions []string, knots int) int {
	visited := make(map[point]struct{})
	knotsList := make([]point, knots)
	lastKnotIndex := len(knotsList) - 1
	visited[knotsList[lastKnotIndex]] = struct{}{}

	for i := range motions {
		split := strings.Fields(motions[i])
		direction := split[0]
		steps, _ := strconv.Atoi(split[1])

		for range steps {
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
						visited[knotsList[k]] = struct{}{}
					}
				}
			}
		}
	}

	return len(visited)
}
