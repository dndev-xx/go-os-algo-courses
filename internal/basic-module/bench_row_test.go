package basicmodule_test

import (
	"math"
	"strconv"
	"testing"

	basicmodule "github.com/dndev-xx/go-os-algo-courses/internal/basic-module" //nolint:depguard
)

// BenchmarkIteractivePow Бенчмарк для IteractivePow.
func BenchmarkIteractivePow(b *testing.B) {
	testCases := []struct {
		base float64
		exp  int
	}{
		{2.0, 10},     // положительная степень
		{2.0, -10},    // отрицательная степень
		{2.5, 20},     // нецелое основание
		{0.5, 15},     // дробное основание
		{3.0, 0},      // нулевая степень
		{0.0, 5},      // нулевое основание
		{1.0001, 100}, // близко к 1
		{0.9999, 100}, // близко к 0
	}

	for _, tc := range testCases {
		b.Run(
			"base="+formatFloat(tc.base)+",exp="+formatInt(tc.exp),
			func(b *testing.B) {
				for b.Loop() {
					_, _ = basicmodule.IteractivePow(tc.base, tc.exp)
				}
			},
		)
	}
}

// BenchmarkBinaryExpansionPow Бенчмарк для IteractivePow.
func BenchmarkBinaryExpansionPow(b *testing.B) {
	testCases := []struct {
		base float64
		exp  int
	}{
		{2.0, 10},     // положительная степень
		{2.0, -10},    // отрицательная степень
		{2.5, 20},     // нецелое основание
		{0.5, 15},     // дробное основание
		{3.0, 0},      // нулевая степень
		{0.0, 5},      // нулевое основание
		{1.0001, 100}, // близко к 1
		{0.9999, 100}, // близко к 0
	}

	for _, tc := range testCases {
		b.Run(
			"base="+formatFloat(tc.base)+",exp="+formatInt(tc.exp),
			func(b *testing.B) {
				for b.Loop() {
					_ = basicmodule.BinaryExpansionPow(tc.base, tc.exp)
				}
			},
		)
	}
}

// BenchmarkMathPow Бенчмарк для math.Pow (адаптер, так как math.Pow принимает float64).
func BenchmarkMathPow(b *testing.B) {
	testCases := []struct {
		base float64
		exp  float64
	}{
		{2.0, 10.0},
		{2.0, -10.0},
		{2.5, 20.0},
		{0.5, 15.0},
		{3.0, 0.0},
		{0.0, 5.0},
		{1.0001, 100.0},
		{0.9999, 100.0},
	}

	for _, tc := range testCases {
		b.Run(
			"base="+formatFloat(tc.base)+",exp="+formatFloat(tc.exp),
			func(b *testing.B) {
				for b.Loop() {
					math.Pow(tc.base, tc.exp)
				}
			},
		)
	}
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func formatInt(i int) string {
	return strconv.Itoa(i)
}

// BenchmarkLargeExponent тест на больших значениях.
func BenchmarkLargeExponent(b *testing.B) {
	b.Run("IteractivePow-large", func(b *testing.B) {
		for b.Loop() {
			_, _ = basicmodule.IteractivePow(1.5, 1000)
		}
	})

	b.Run("MathPow-large", func(b *testing.B) {
		for b.Loop() {
			math.Pow(1.5, 1000.0)
		}
	})

	b.Run("BinaryExpansionPow", func(b *testing.B) {
		for b.Loop() {
			_ = basicmodule.BinaryExpansionPow(1.5, 1000)
		}
	})
}
