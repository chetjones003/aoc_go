package myMath

// Takes in an int slice: [1 2 3].
// Returns the sum: 6.
func SumSlice(measurement []int) int {
	sum := 0

	for _, value := range measurement {
		sum += value
	}

	return sum
}
