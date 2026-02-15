package basicmodule_test

import (
	"strconv"
	"testing"

	basicmodule "github.com/dndev-xx/go-os-algo-courses/internal/basic-module" //nolint:depguard
)

// BenchmarkIteractivePow Бенчмарк для IterectiveFibonacci.
func BenchmarkIterectiveFibonacci(b *testing.B) {
	testCases := []struct {
		base int
	}{
		{base: 5},
		{base: 10},
		{base: 15},
		{base: 20},
		{base: 25},
		{base: 30},
	}

	for _, tc := range testCases {
		b.Run(
			"base="+strconv.Itoa(tc.base),
			func(b *testing.B) {
				for b.Loop() {
					_ = basicmodule.Fibonacci(tc.base)
				}
			},
		)
	}
}

// BenchmarkIteractivePow Бенчмарк для IterectiveFibonacci.
func BenchmarkRecFibonacci(b *testing.B) {
	testCases := []struct {
		base int
	}{
		{base: 5},
		{base: 10},
		{base: 15},
		{base: 20},
		{base: 25},
		{base: 30},
	}

	for _, tc := range testCases {
		b.Run(
			"base="+strconv.Itoa(tc.base),
			func(b *testing.B) {
				for b.Loop() {
					_ = basicmodule.RecFibonacci(tc.base)
				}
			},
		)
	}
}

// BenchmarkIteractivePow Бенчмарк для IterectiveFibonacci.
func BenchmarkGoldenRationFibonacci(b *testing.B) {
	testCases := []struct {
		base float64
	}{
		{base: 5},
		{base: 10},
		{base: 15},
		{base: 20},
		{base: 25},
		{base: 30},
	}

	for _, tc := range testCases {
		b.Run(
			"base="+formatFloat(tc.base),
			func(b *testing.B) {
				for b.Loop() {
					_ = basicmodule.GoldenRatioFibonacci(tc.base)
				}
			},
		)
	}
}
