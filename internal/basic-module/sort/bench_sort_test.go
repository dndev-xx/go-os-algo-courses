//nolint:depguard,gosec
package sort_test

import (
	"math/rand/v2"
	"testing"

	"github.com/dndev-xx/go-os-algo-courses/internal/basic-module/sort"
)

func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range size {
		slice[i] = rand.IntN(1000)
	}
	return slice
}

func BenchmarkBubbleSort100(b *testing.B) {
	arr := generateRandomSlice(100)
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		b.StartTimer()

		sort.BubbleSort(tmp, shouldIntSwap)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	arr := generateRandomSlice(1000)
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		b.StartTimer()

		sort.BubbleSort(tmp, shouldIntSwap)
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	arr := generateRandomSlice(10000)
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		b.StartTimer()

		sort.BubbleSort(tmp, shouldIntSwap)
	}
}

func BenchmarkInsertionSort100(b *testing.B) {
	arr := generateRandomSlice(100)
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		b.StartTimer()

		sort.InsertionSort(tmp, shouldIntSwap)
	}
}

func BenchmarkInsertionSort1000(b *testing.B) {
	arr := generateRandomSlice(1000)
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		b.StartTimer()

		sort.InsertionSort(tmp, shouldIntSwap)
	}
}

func BenchmarkInsertionSort10000(b *testing.B) {
	arr := generateRandomSlice(10000)
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		b.StartTimer()

		sort.InsertionSort(tmp, shouldIntSwap)
	}
}

func BenchmarkShellSort100(b *testing.B) {
	arr := generateRandomSlice(100)
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		b.StartTimer()

		sort.ShellSort(tmp, shouldIntSwap)
	}
}

func BenchmarkShellSort1000(b *testing.B) {
	arr := generateRandomSlice(1000)
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		b.StartTimer()

		sort.ShellSort(tmp, shouldIntSwap)
	}
}

func BenchmarkShellSort10000(b *testing.B) {
	arr := generateRandomSlice(10000)
	b.ResetTimer()

	for b.Loop() {
		b.StopTimer()
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		b.StartTimer()

		sort.ShellSort(tmp, shouldIntSwap)
	}
}

func BenchmarkCompareSorts(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		baseArr := generateRandomSlice(size)

		b.Run("BubbleSort/"+itoa(size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				arr := make([]int, len(baseArr))
				copy(arr, baseArr)
				b.StartTimer()

				sort.BubbleSort(arr, shouldIntSwap)
			}
		})

		b.Run("InsertionSort/"+itoa(size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				arr := make([]int, len(baseArr))
				copy(arr, baseArr)
				b.StartTimer()

				sort.InsertionSort(arr, shouldIntSwap)
			}
		})

		b.Run("ShellSort/"+itoa(size), func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				arr := make([]int, len(baseArr))
				copy(arr, baseArr)
				b.StartTimer()

				sort.ShellSort(arr, shouldIntSwap)
			}
		})
	}
}

func itoa(n int) string {
	if n == 100 {
		return "100"
	}
	if n == 1000 {
		return "1000"
	}
	return "10000"
}

func shouldIntSwap(f, s int) bool {
	return f < s
}

func shouldShellIntSwap(f, s int) bool {
	return f <= s
}

func shouldInsertionSwap(f, s int) bool {
	return f >= s
}
