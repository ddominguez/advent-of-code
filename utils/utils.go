package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Puzzle struct {
	Day        string
	Year       string
	UseExample bool
}

// rootPath uses git to return the root path of the project
func rootPath() (string, error) {
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

// inputFile returns the path of the input file
func (p Puzzle) inputFile() (string, error) {
	root, err := rootPath()
	if err != nil {
		return "", err
	}

	var fn = fmt.Sprintf("%s.txt", p.Day)
	if p.UseExample {
		fn = fmt.Sprintf("%s.example.txt", p.Day)
	}

	fp := path.Join(root, p.Year, "input", fn)
	if _, err := os.Stat(fp); err != nil {
		return "", fmt.Errorf("%s does not exist", fp)
	}

	return fp, nil
}

// InputData will read the input file and return the contents of the file
func InputData(p Puzzle) (string, error) {
	filePath, err := p.inputFile()
	if err != nil {
		return "", err
	}
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(dat)), nil
}

func InputDataNoTrim(p Puzzle) (string, error) {
	filePath, err := p.inputFile()
	if err != nil {
		return "", err
	}
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}
