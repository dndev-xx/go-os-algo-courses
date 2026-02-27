//nolint:gosec,errcheck,depguard,thelper
package slices_test

import (
	"fmt"
	"testing"

	"github.com/dndev-xx/go-os-algo-courses/internal/basic-module/slices"
)

func BenchmarkDArray_Add_Small(b *testing.B) {
	sizes := []int{10, 100}
	strategies := []struct {
		name string
		st   slices.LoadType
	}{
		{"AddOne", slices.AddOne},
		{"AddHundred", slices.AddHundred},
		{"MultiplyByTwo", slices.MultiplyByTwo},
	}

	for _, size := range sizes {
		for _, strategy := range strategies {
			b.Run(fmt.Sprintf("%s_%d", strategy.name, size), func(b *testing.B) {
				benchmarkDArrayAdd(b, strategy.st, size)
			})
		}
	}
}

func Benchmark_GoSlice_Small(b *testing.B) {
	sizes := []int{10, 100}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("GoSlice_%d", size), func(b *testing.B) {
			benchmarkGoSliceAdd(b, size)
		})
	}
}

func Benchmark_InsertMiddle_Small(b *testing.B) {
	sizes := []int{10, 50}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("DArray_MultiplyByTwo_%d", size), func(b *testing.B) {
			benchmarkDArrayInsertMiddle(b, slices.MultiplyByTwo, size)
		})

		b.Run(fmt.Sprintf("GoSlice_%d", size), func(b *testing.B) {
			benchmarkGoSliceInsertMiddle(b, size)
		})
	}
}

func Benchmark_GrowStrategies_Comparison(b *testing.B) {
	size := 100
	strategies := []struct {
		name string
		st   slices.LoadType
	}{
		{"AddOne", slices.AddOne},
		{"AddHundred", slices.AddHundred},
		{"MultiplyByTwo", slices.MultiplyByTwo},
	}

	for _, strategy := range strategies {
		b.Run(strategy.name, func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				da := slices.New[int](0, 0, strategy.st)
				b.StartTimer()

				for j := range size {
					_ = da.Add(j, j)
				}

				b.StopTimer()
				da.Close()
			}
		})
	}

	b.Run("GoSlice", func(b *testing.B) {
		for b.Loop() {
			slice := make([]int, 0)
			for j := range size {
				_ = append(slice, j)
			}
		}
	})
}

func Benchmark_GrowOverhead_Small(b *testing.B) {
	size := 50
	strategies := []struct {
		name string
		st   slices.LoadType
	}{
		{"AddOne", slices.AddOne},
		{"MultiplyByTwo", slices.MultiplyByTwo},
	}

	for _, strategy := range strategies {
		b.Run("DArray_"+strategy.name, func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				da := slices.New[int](0, 1, strategy.st)
				b.StartTimer()

				for j := range size {
					_ = da.Add(j, j)
				}

				b.StopTimer()
				da.Close()
			}
		})
	}

	b.Run("GoSlice", func(b *testing.B) {
		for b.Loop() {
			slice := make([]int, 0, 1)
			for j := range size {
				_ = append(slice, j)
			}
		}
	})
}

func benchmarkDArrayAdd(b *testing.B, strategy slices.LoadType, size int) {
	b.ResetTimer()
	for b.Loop() {
		b.StopTimer()
		da := slices.New[int](0, 0, strategy)
		b.StartTimer()

		for j := range size {
			_ = da.Add(j, j)
		}

		b.StopTimer()
		da.Close()
	}
}

func benchmarkGoSliceAdd(b *testing.B, size int) {
	b.ResetTimer()
	for b.Loop() {
		slice := make([]int, 0)
		for j := range size {
			_ = append(slice, j)
		}
	}
}

func benchmarkDArrayInsertMiddle(b *testing.B, strategy slices.LoadType, size int) {
	da := slices.New[int](0, size, strategy)
	for j := range size {
		_ = da.Add(j, j)
	}
	defer da.Close()

	b.ResetTimer()
	for b.Loop() {
		mid := size / 2
		_ = da.Add(999, mid)
		_ = da.Remove(mid)
	}
}

func benchmarkGoSliceInsertMiddle(b *testing.B, size int) {
	slice := make([]int, size)
	for j := range size {
		slice[j] = j
	}

	b.ResetTimer()
	for b.Loop() {
		mid := size / 2
		slice = append(slice[:mid], append([]int{999}, slice[mid:]...)...)
		slice = append(slice[:mid], slice[mid+1:]...)
	}
}
