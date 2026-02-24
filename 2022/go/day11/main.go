package main

import (
	"container/list"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "")
	flag.Parse()
	data, _ := os.ReadFile("../../input/11.txt")

	if *part == 2 {
		fmt.Println(part2(string(data)))
	} else {
		fmt.Println(part1(string(data)))
	}
}

type Monkey struct {
	inspections   int
	worryLevels   *list.List
	opType        string
	opVal         int
	useOld        bool
	divisibleBy   int
	testOkMonkey  int
	testBadMonkey int
}

func newMonkey(lines []string) *Monkey {
	m := &Monkey{worryLevels: list.New()}
	for i := range lines {
		ln := strings.TrimSpace(lines[i])
		switch {
		case strings.HasPrefix(ln, "Starting items:"):
			after := strings.TrimPrefix(ln, "Starting items:")
			for _, l := range strings.Fields(after) {
				parsed, _ := strconv.Atoi(strings.TrimRight(l, ","))
				m.worryLevels.PushBack(parsed)
			}
		case strings.HasPrefix(ln, "Operation:"):
			after := strings.TrimPrefix(ln, "Operation: new =")
			f := strings.Fields(after)
			m.opType = f[1]
			if f[2] == "old" {
				m.useOld = true
			} else {
				parsed, _ := strconv.Atoi(f[2])
				m.opVal = parsed
			}
		case strings.HasPrefix(ln, "Test:"):
			after := strings.TrimPrefix(ln, "Test: divisible by ")
			parsed, _ := strconv.Atoi(after)
			m.divisibleBy = parsed
		case strings.HasPrefix(ln, "If true:"):
			after := strings.TrimPrefix(ln, "If true: throw to monkey ")
			parsed, _ := strconv.Atoi(after)
			m.testOkMonkey = parsed
		case strings.HasPrefix(ln, "If false:"):
			after := strings.TrimPrefix(ln, "If false: throw to monkey ")
			parsed, _ := strconv.Atoi(after)
			m.testBadMonkey = parsed
		}
	}

	return m
}

func (m Monkey) testPasses(worryLevel int) bool {
	return worryLevel%m.divisibleBy == 0
}

func (m Monkey) doOperation(worryLevel int) int {
	var opVal int
	if m.useOld {
		opVal = worryLevel
	} else {
		opVal = m.opVal
	}

	switch m.opType {
	case "*":
		worryLevel = worryLevel * opVal
	case "+":
		worryLevel = worryLevel + opVal
	}

	return worryLevel
}

func makeMonkeyMap(items []string) map[int]*Monkey {
	mm := make(map[int]*Monkey)
	for i := range items {
		mm[i] = newMonkey(strings.Split(items[i], "\n"))
	}
	return mm
}

func getMonkeyBusiness(monkeyMap map[int]*Monkey) int {
	var a, b int
	for _, v := range monkeyMap {
		if v.inspections > a {
			b = a
			a = v.inspections
		} else if v.inspections > b {
			b = v.inspections
		}
	}
	return a * b
}

func part1(input string) int {
	mm := makeMonkeyMap(strings.Split(input, "\n\n"))
	rounds := 20
	withRelief := true
	mm = doInspections(mm, rounds, withRelief)
	res := getMonkeyBusiness(mm)
	return res
}

func part2(input string) int {
	mm := makeMonkeyMap(strings.Split(input, "\n\n"))
	rounds := 10_000
	withRelief := false
	mm = doInspections(mm, rounds, withRelief)
	res := getMonkeyBusiness(mm)
	return res
}

func doInspections(monkeyMap map[int]*Monkey, rounds int, withRelief bool) map[int]*Monkey {
	var prodDivisible int
	reliefVal := 3
	if !withRelief {
		prodDivisible = getProductOfDivisible(monkeyMap)
	}
	for range rounds {
		for i := range len(monkeyMap) {
			m := monkeyMap[i]
			for m.worryLevels.Len() > 0 {
				m.inspections++
				f := m.worryLevels.Front()
				worryLevel := m.worryLevels.Remove(f).(int)
				worryLevel = m.doOperation(worryLevel)
				if withRelief {
					worryLevel = worryLevel / reliefVal
				} else {
					worryLevel = worryLevel % prodDivisible
				}
				if m.testPasses(worryLevel) {
					monkeyMap[m.testOkMonkey].worryLevels.PushBack(worryLevel)
				} else {
					monkeyMap[m.testBadMonkey].worryLevels.PushBack(worryLevel)
				}
			}
			monkeyMap[i] = m
		}
	}

	return monkeyMap
}

func getProductOfDivisible(monkeyMap map[int]*Monkey) int {
	product := 1
	for _, m := range monkeyMap {
		product = product * m.divisibleBy
	}
	return product
}
