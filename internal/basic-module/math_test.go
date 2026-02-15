package basicmodule_test

import (
	"testing"

	basicmodule "github.com/dndev-xx/go-os-algo-courses/internal/basic-module" //nolint:depguard
	"github.com/stretchr/testify/assert"                                       //nolint:depguard
	"github.com/stretchr/testify/require"                                      //nolint:depguard
)

func TestIteractivePow(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n       float64
		pow     int
		want    float64
		wantErr bool
	}{
		{
			name:    "test_simple_pow",
			n:       2.0,
			pow:     3,
			want:    8.0,
			wantErr: false,
		},
		{
			name:    "test_simple_pow_negative_number",
			n:       2.0,
			pow:     -3,
			want:    0.125,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := basicmodule.IteractivePow(tt.n, tt.pow)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("IteractivePow() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("IteractivePow() succeeded unexpectedly")
			}
			require.NoError(t, gotErr)
			assert.InEpsilon(t, tt.want, got, 0.000001)
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{
			name: "test01_simple_fib",
			n:    5,
			want: 5,
		},
		{
			name: "test02_simple_fib",
			n:    10,
			want: 55,
		},
		{
			name: "test03_simple_fib",
			n:    41,
			want: 165580141,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basicmodule.Fibonacci(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRecFibonacci(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{
			name: "test01_simple_fib",
			n:    5,
			want: 5,
		},
		{
			name: "test02_simple_fib",
			n:    10,
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basicmodule.RecFibonacci(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFindSimpleDiv(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want []int
	}{
		{
			name: "test01",
			n:    50,
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
		{
			name: "test02",
			n:    100,
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basicmodule.FindSimpleDiv(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBinaryExpansionPow(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    float64
		pow  int
		want float64
	}{
		{
			name: "test01",
			n:    2.0,
			pow:  3,
			want: 8.0,
		},
		{
			name: "test02",
			n:    5.0,
			pow:  3,
			want: 125,
		},
		{
			name: "test03",
			n:    2.0,
			pow:  -3,
			want: 0.125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basicmodule.BinaryExpansionPow(tt.n, tt.pow)
			assert.InEpsilon(t, tt.want, got, 0.000001)
		})
	}
}

func TestGoldenRatioFibonacci(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    float64
		want uint64
	}{
		{
			name: "test01_simple_fib",
			n:    5,
			want: 5,
		},
		{
			name: "test02_simple_fib",
			n:    10,
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basicmodule.GoldenRatioFibonacci(tt.n)
			assert.InEpsilon(t, tt.want, got, 0.000001)
		})
	}
}
