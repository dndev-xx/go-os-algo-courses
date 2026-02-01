// Package basicmodule provides a basic module for testing.
// Contains algorithms for generating 25x25 grid patterns.
package basicmodule

import "strings"

// Constants defining grid dimensions.
const (
	X = 25 // Grid width (number of columns)
	Y = 25 // Grid height (number of rows)
)

// Gpmagic generates a 25x25 grid based on a condition function.
// Each cell is represented by '#' if condition(i, j) is true, '.' otherwise.
// Returns the grid as a string with newline separators and an error (always nil currently).
func Gpmagic(cond func(i, j int) bool) (string, error) {
	var (
		rsl strings.Builder
		err error
	)
	for i := range X {
		for j := range Y {
			if cond(i, j) {
				rsl.WriteString("#")
			} else {
				rsl.WriteString(".")
			}
		}
		rsl.WriteString("\n")
	}
	return rsl.String(), err
}

// IsDiagonalCond returns true if i > j.
// Creates a diagonal pattern where cells above the main diagonal are filled.
func IsDiagonalCond(i, j int) bool {
	return i < j
}

// IsStrongDiagonalCond return true if i == j.
// Creates a diagonal pattern where '#' strong diagonal is line.
func IsStrongDiagonalCond(i, j int) bool {
	return i == j
}
