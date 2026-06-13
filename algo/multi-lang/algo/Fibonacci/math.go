package Fibonacci

import (
	"math"
)

func Round(f float64) float64 {
	mod := math.Mod(10*(f-math.Floor(f)), 10)
	if mod <= 4 {
		return math.Floor(f)
	}
	return math.Ceil(f)
}

func Fibonacci_math(n float64) float64 {
	A := math.Sqrt(5) / 5
	B := (1 + math.Sqrt(5)) / 2
	C := (1 - math.Sqrt(5)) / 2
	return Round(A*math.Pow(B, n) - math.Pow(C, n))
}
