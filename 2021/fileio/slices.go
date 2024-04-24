package fileio

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLinesAndSplitByDelimeter(path string, delimeter string) [][]string {
	var values [][]string
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		values = append(values, strings.Split(scanner.Text(), delimeter))
	}

	defer file.Close()
	return values
}

func ReadLinesIntoIntSlice(path string) []int {
	var values []int
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}

	defer file.Close()
	return values
}
