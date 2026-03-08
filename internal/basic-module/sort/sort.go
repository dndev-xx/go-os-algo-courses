// Package sort impl all popular sorting algorithms
//
//nolint:revive
package sort

func BubbleSort[T comparable](arr []T, shouldSwap func(f, s T) bool) {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if shouldSwap(arr[j], arr[j-1]) {
				swap(arr, j, j-1)
			}
		}
	}
}

func OptBubbleSort[T comparable](arr []T, shouldSwap func(f, s T) bool) {
	var isSwap bool
	for i := 0; i < len(arr)-1; i++ {
		isSwap = false
		for j := len(arr) - 1; j > i; j-- {
			if shouldSwap(arr[j], arr[j-1]) {
				isSwap = true
				swap(arr, j, j-1)
			}
		}
		if !isSwap {
			break
		}
	}
}

func InsertionSort[T comparable](arr []T, shouldSwap func(f, s T) bool) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && shouldSwap(arr[j], arr[j-1]); j-- {
			swap(arr, j, j-1)
		}
	}
}

func OptInsertionSort[T comparable](arr []T, shouldSwap func(f, s T) bool) {
	for i := 1; i < len(arr); i++ {
		index := i
		cur := arr[i]
		for index > 0 {
			if shouldSwap(cur, arr[index-1]) {
				break
			}
			arr[index] = arr[index-1]
			index--
		}
		arr[index] = cur
	}
}

func OptInsertionSortByBinarySearch(arr []int) {
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		pos := binarySearch(arr, cur, 0, i-1)
		for j := i - 1; j >= pos; j-- {
			arr[j+1] = arr[j]
		}
		arr[pos] = cur
	}
}

func ShellSort[T comparable](arr []T, shouldSwap func(f, s T) bool) {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			for j := i; j >= gap && shouldSwap(arr[j], arr[j-gap]); j -= gap {
				swap(arr, j, j-gap)
			}
		}
	}
}

func OptShellSort[T comparable](arr []T, shouldSwap func(f, s T) bool) {
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			for j := i; j >= gap; j -= gap {
				if shouldSwap(arr[j-gap], arr[j]) {
					break
				}
				swap(arr, j-gap, j)
			}
		}
	}
}

func swap[T comparable](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func binarySearch(arr []int, key, left, right int) int {
	if left > right {
		return left
	}
	mid := (left + right) / 2
	if key < arr[mid] {
		return binarySearch(arr, key, left, mid-1)
	}
	return binarySearch(arr, key, mid+1, right)
}
