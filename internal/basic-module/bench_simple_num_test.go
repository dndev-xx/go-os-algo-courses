package basicmodule_test

import (
	"strconv"
	"testing"

	basicmodule "github.com/dndev-xx/go-os-algo-courses/internal/basic-module" //nolint:depguard
)

// BenchmarkIteractivePow Бенчмарк для FindSimpleDiv.
func BenchmarkFindSimpleDiv(b *testing.B) {
	testCases := []struct {
		base int
	}{
		{base: 100},
		{base: 1000},
		{base: 10000},
		{base: 100000},
		// long time {base: 1000000},
		// long time {base: 10000000},
	}

	for _, tc := range testCases {
		b.Run(
			"base="+strconv.Itoa(tc.base),
			func(b *testing.B) {
				for b.Loop() {
					_ = basicmodule.FindSimpleDiv(tc.base)
				}
			},
		)
	}
}

// BenchmarkIteractivePow Бенчмарк для FindDiv.
func BenchmarkFindDiv(b *testing.B) {
	testCases := []struct {
		base int
	}{
		{base: 100},
		{base: 1000},
		{base: 10000},
		{base: 100000},
		{base: 1000000},
		{base: 10000000},
	}

	for _, tc := range testCases {
		b.Run(
			"base="+strconv.Itoa(tc.base),
			func(b *testing.B) {
				for b.Loop() {
					_ = basicmodule.FindDiv(tc.base)
				}
			},
		)
	}
}

// BenchmarkIteractivePow Бенчмарк для FindEratosthenes.
func BenchmarkFindEratosthenes(b *testing.B) {
	testCases := []struct {
		base int
	}{
		{base: 100},
		{base: 1000},
		{base: 10000},
		{base: 100000},
		// long time {base: 1000000},
		// long time {base: 10000000},
	}

	for _, tc := range testCases {
		b.Run(
			"base="+strconv.Itoa(tc.base),
			func(b *testing.B) {
				for b.Loop() {
					_ = basicmodule.FindEratosthenes(tc.base)
				}
			},
		)
	}
}
