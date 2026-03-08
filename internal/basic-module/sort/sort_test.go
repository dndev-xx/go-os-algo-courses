//nolint:depguard
package sort_test

import (
	"testing"

	"github.com/dndev-xx/go-os-algo-courses/internal/basic-module/sort"
	"github.com/stretchr/testify/assert"
)

func Test_CommonSorting(t *testing.T) {
	type testCase[T, F comparable] struct {
		name       string
		arr        []T
		sortFunc   func([]T, func(T, F) bool)
		shouldSwap func(T, F) bool
		expected   []T
	}

	tests := []testCase[int, int]{
		{
			name:       "BubbleSort integers small array",
			arr:        []int{3, 1, 4, 1, 5},
			sortFunc:   sort.BubbleSort[int],
			shouldSwap: shouldIntSwap,
			expected:   []int{1, 1, 3, 4, 5},
		},
		{
			name:       "OptBubbleSort integers small array",
			arr:        []int{3, 1, 4, 1, 5},
			sortFunc:   sort.OptBubbleSort[int],
			shouldSwap: shouldIntSwap,
			expected:   []int{1, 1, 3, 4, 5},
		},
		{
			name:       "InsertionSort integers small array",
			arr:        []int{3, 1, 4, 1, 5},
			sortFunc:   sort.InsertionSort[int],
			shouldSwap: shouldIntSwap,
			expected:   []int{1, 1, 3, 4, 5},
		},
		{
			name:       "OptInsertionSort integers small array",
			arr:        []int{3, 1, 4, 1, 5},
			sortFunc:   sort.OptInsertionSort[int],
			shouldSwap: shouldInsertionSwap,
			expected:   []int{1, 1, 3, 4, 5},
		},
		{
			name:       "ShellSort integers small array",
			arr:        []int{3, 1, 4, 1, 5},
			sortFunc:   sort.ShellSort[int],
			shouldSwap: shouldIntSwap,
			expected:   []int{1, 1, 3, 4, 5},
		},
		{
			name:       "OptShellSort integers small array",
			arr:        []int{3, 1, 4, 1, 5},
			sortFunc:   sort.OptShellSort[int],
			shouldSwap: shouldShellIntSwap,
			expected:   []int{1, 1, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sortFunc(tt.arr, tt.shouldSwap)
			assert.Equal(t, tt.expected, tt.arr)
		})
	}
}

func Test_InsertionSortByBinarySearch(t *testing.T) {
	t.Run("simple tets", func(t *testing.T) {
		arr := []int{3, 1, 4, 1, 5}
		sort.OptInsertionSortByBinarySearch(arr)
		assert.Equal(t, []int{1, 1, 3, 4, 5}, arr)
	})
}
