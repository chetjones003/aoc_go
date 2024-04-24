package myMath

import (
	"testing"
)

func TestSumSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			"test 1",
			[]int{1, 2, 3},
			6,
		},
		{
			"test 2",
			[]int{100, 25, 2256},
			2381,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumSlice(tt.input)
			if result != tt.want {
				t.Errorf("got: %d, expected: %d", result, tt.want)
			}
		})
	}
}
