package basicmodule

import "math"

// IteractivePow func is implementation of math.Pow
// Iterative algorithm O(N) time complexity
// Param:
// n - base number
// pow - base pow may negative
// return float64.
func IteractivePow(n float64, pow int) (float64, error) {
	flNegative := false
	src := 1.0
	if n == 0 {
		return src, nil
	}
	if pow < 0 {
		pow = -pow
		flNegative = true
	}
	for range pow {
		src *= n
	}
	if flNegative {
		src = 1 / src
	}
	return src, nil
}

// BinaryExpansionPow algorithm O(LogN) time complexity.
// x^y = x^y/2 * x^y/2 		where y - even
// x^y = x^y/2 * x^y/2 * x 	where y - odd.
func BinaryExpansionPow(n float64, pow int) float64 {
	isNegative := false
	res := 1.0
	if pow == 0 {
		return res
	}
	if pow < 0 {
		pow = -pow
		isNegative = true
	}
	r := BinaryExpansionPow(n, pow/2)
	if pow%2 != 0 {
		res = n * r * r
	} else {
		res = r * r
	}
	if isNegative {
		res = 1 / res
	}
	return res
}

// RecFibonacci algorithm O(2^N) time complexity.
func RecFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return RecFibonacci(n-1) + RecFibonacci(n-2)
}

// Fibonacci algorithm O(N) time complexity.
func Fibonacci(n int) int {
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}

// GoldenRatioFibonacci algorithm O(N) time complexity
// fi = (1 + sqrt(5))/2
// Fn = fi^n / sqrt(5) + 1/2.
func GoldenRatioFibonacci(n float64) uint64 {
	fi := (1.0 + math.Sqrt(5)) / 2.0
	return uint64(math.Floor(math.Pow(fi, n)/math.Sqrt(5) + 1.0/2.0))
}

// FindSimpleDiv algorithm O(N^2) time complexity.
func FindSimpleDiv(n int) []int {
	rsl := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		cnt := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				cnt++
			}
			if cnt > 2 {
				break
			}
		}
		if cnt == 2 {
			rsl = append(rsl, i)
		}
	}
	return rsl
}
