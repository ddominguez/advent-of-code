package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
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

// InputFile returns the path of the input file
func InputFile(day, year string) (string, error) {
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

// InputData will read the input file and return the contents of the file
func InputData(filePath string) (string, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(dat)), nil
}
