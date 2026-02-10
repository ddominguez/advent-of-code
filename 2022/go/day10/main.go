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
	data, _ := os.ReadFile("../../input/10.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

type Instruction struct {
	addx   int
	cycles int
}

func newInstruction(s string) Instruction {
	if s == "noop" {
		return Instruction{cycles: 1}
	}
	split := strings.Split(s, " ")
	v, _ := strconv.Atoi(split[1])
	return Instruction{addx: v, cycles: 2}
}

func part1(input string) int {
	var signalStrength int
	lines := strings.Split(strings.TrimSpace(input), "\n")
	initialCycleCheck := 20
	cycleInterval := 40
	cycleCheck := initialCycleCheck
	x := 1
	cycle := 0
	for i := range lines {
		instruction := newInstruction(lines[i])

		for c := range instruction.cycles {
			cycle += 1
			if cycle == cycleCheck {
				signalStrength += cycle * x
				cycleCheck += cycleInterval
			}

			if c+1 == instruction.cycles {
				x += instruction.addx
			}
		}
	}
	return signalStrength
}

func part2(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var crtRowSB strings.Builder
	crtRows := []string{}
	crtWidth := 40
	pixelLit := "#"
	pixelDark := "."
	pixelPos := 0
	x := 1
	for i := range lines {
		instruction := newInstruction(lines[i])

		for c := range instruction.cycles {
			if pixelPos == x-1 || pixelPos == x || pixelPos == x+1 {
				crtRowSB.WriteString(pixelLit)
			} else {
				crtRowSB.WriteString(pixelDark)
			}

			pixelPos += 1
			if pixelPos == crtWidth {
				crtRows = append(crtRows, crtRowSB.String())
				crtRowSB.Reset()
				pixelPos = 0
			}

			if c+1 == instruction.cycles {
				x += instruction.addx
			}
		}
	}

	return strings.Join(crtRows, "\n")
}
