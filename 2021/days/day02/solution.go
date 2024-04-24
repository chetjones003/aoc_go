package day02

import (
	"aoc/fileio"
	"fmt"
	"strconv"
)

func Solution(path string) {
	values := fileio.ReadLinesAndSplitByDelimeter(path, " ")

	fmt.Printf("Part 1: %v\n", part1(values))
	fmt.Printf("Part 2: %v\n", part2(values))
}

func part1(instructions [][]string) int {
	position := []int{0, 0}
	for _, instruction := range instructions {
		position = applyMovement(position, instruction)
	}

	result := position[0] * position[1]
	return result
}

func part2(instructions [][]string) int {
	position := []int{0, 0}
	aim := 0

	for _, instruction := range instructions {
		position, aim = applyAim(position, aim, instruction)
	}

	result := position[0] * position[1]
	return result
}

func applyMovement(position []int, instruction []string) []int {
	moveValue, _ := strconv.Atoi(instruction[1])
	var newPosition []int

	switch instruction[0] {
	case "forward":
		newPosition = []int{position[0] + moveValue, position[1]}
	case "down":
		newPosition = []int{position[0], position[1] + moveValue}
	case "up":
		newPosition = []int{position[0], position[1] - moveValue}
	}

	return newPosition
}

func applyAim(position []int, aim int, instruction []string) ([]int, int) {
	changeValue, _ := strconv.Atoi(instruction[1])
	var newPosition = position
	var newAim = aim

	switch instruction[0] {
	case "forward":
		newPosition = []int{position[0] + changeValue, position[1] + changeValue*aim}
	case "down":
		newAim += changeValue
	case "up":
		newAim -= changeValue
	}
	return newPosition, newAim
}
