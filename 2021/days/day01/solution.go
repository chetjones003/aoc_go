package day01

import (
	"aoc/fileio"
	"aoc/myMath"
	"fmt"
)

func Solution(path string) {
	values := fileio.ReadLinesIntoIntSlice(path)
	part1 := calcDepthIncrease(values, 1)
	part2 := calcDepthIncrease(values, 3)

	fmt.Printf("Part1: %v\n", part1)
	fmt.Printf("Part2: %v\n", part2)
}

func calcDepthIncrease(values []int, windowSize int) int {
	numValues := len(values)
	count := 0

	for i := 0; i < numValues-windowSize; i++ {
		window1 := values[i : i+windowSize]
		window2 := values[i+1 : i+1+windowSize]

		if myMath.SumSlice(window2) > myMath.SumSlice(window1) {
			count++
		}
	}

	return count
}
