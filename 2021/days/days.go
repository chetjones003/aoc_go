package days

import (
	"aoc/days/day01"
	"aoc/days/day02"
	"fmt"
	"time"
)

type dayFunctionType func(string)

func Run(day string, path string) {
	startTime := time.Now()

	dayFunctions := map[string]dayFunctionType{
		"01": day01.Solution,
		"02": day02.Solution,
	}
	dayFunction, found := dayFunctions[day]
	if found {
		dayFunction(path)
	} else {
		fmt.Printf("No code ready for day %s\n", day)
	}

	fmt.Println(time.Since(startTime))
	return
}
