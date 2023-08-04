package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"sort"
	"strconv"
	"strings"
)

// rootPath uses git to return the root path of the project
func rootPath() (string, error) {
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

// inputFile returns the path of the input file
func inputFile(day, year string) (string, error) {
	root, err := rootPath()
	if err != nil {
		return "", err
	}

	fp := path.Join(root, year, "input", fmt.Sprintf("%s.txt", day))
	if _, err := os.Stat(fp); err != nil {
		return "", fmt.Errorf("%s does not exist", fp)
	}

	return fp, nil
}

// inputData will read the input file and return the contents of the file
func inputData(filePath string) (string, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(dat)), nil
}

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
	day := "01"
	year := "2022"

	inputFile, err := inputFile(day, year)
	if err != nil {
		panic(err)
	}

	data, err := inputData(inputFile)
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
