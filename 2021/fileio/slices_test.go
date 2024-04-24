package fileio

import (
	"fmt"
	"testing"
)

func TestReadLinesAndSplitByDelimeter(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		delimeter string
		want      [][]string
	}{
		{"can slice by ' '", "test_inputs/test_input_slice_delimeter1.txt", " ", [][]string{
			{"forward", "5"},
			{"down", "5"},
			{"forward", "8"},
			{"up", "3"},
			{"down", "8"},
			{"forward", "2"},
		},
		},
		{"can slice by ','", "test_inputs/test_input_slice_delimeter2.txt", ",", [][]string{
			{"forward", "5"},
			{"down", "5"},
			{"forward", "8"},
			{"up", "3"},
			{"down", "8"},
			{"forward", "2"},
		},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReadLinesAndSplitByDelimeter(tt.input, tt.delimeter)

			if len(result) != len(tt.want) {
				t.Errorf("Lengths do not match: got: %d, want: %d", len(result), len(tt.want))
			}

			for i, slice := range result {
				for j := range slice {
					if result[i][j] != tt.want[i][j] {
						t.Errorf("got: %s, expected: %s", result[i][j], tt.want[i][j])
						fmt.Println("Result:", result)
						fmt.Println("tt.want:", tt.want)
					}
				}
			}
		})
	}
}

func TestReadLinesIntoIntSlice(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{"can return a slice of integers", "test_inputs/test_input_slice_of_ints.txt", []int{
			1,
			2,
			3,
			4,
			5,
			6,
			7,
			8,
			9,
			10,
			1000000,
			200000,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReadLinesIntoIntSlice(tt.input)

			for i, value := range result {
				if value != tt.want[i] {
					t.Errorf("got: %d, expected: %d", value, tt.want[i])
					fmt.Println("Result:", result)
					fmt.Println("tt.want:", tt.want)
				}
			}
		})
	}
}
