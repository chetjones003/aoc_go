package main

import (
	"aoc/days"
	"fmt"
	"os"
	"time"
)

func main() {
	var day, inputFilePrefix string

	args := os.Args[1:]

	inputFilePrefix = args[0]
	if len(args) > 1 {
		day = args[1]
	} else {
		// Use todays date to choose which problem to solve
		day = fmt.Sprintf("%02d", time.Now().Day())
	}

	inputPath := fmt.Sprintf("inputs/%s/%s.txt", day, inputFilePrefix)
	fmt.Printf("Running AOC for day %s with input '%s'\n", day, inputPath)
	days.Run(day, inputPath)
}
