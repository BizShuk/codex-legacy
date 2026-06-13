package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%f\n", Sqrt(4))
	fmt.Printf("%f\n", Sqrt(5))
	fmt.Printf("%f\n", Sqrt(9))
	fmt.Printf("%f\n", Sqrt(16))
	fmt.Printf("%f\n", math.Sqrt(4))
	fmt.Printf("%f\n", math.Sqrt(5))
	fmt.Printf("%f\n", math.Sqrt(9))
	fmt.Printf("%f\n", math.Sqrt(16))

	fmt.Printf("%f\n", Round(1.7, 0))
	fmt.Printf("%f\n", Round(1.17, -1))
	fmt.Printf("%f\n", Round(1.119, -2))
	fmt.Printf("%f\n", Round(1.1114, -3))
	fmt.Printf("%f\n", Round(15.11, 1))
	fmt.Printf("%f\n", Round(141.11, 2))
	fmt.Printf("%f\n", Round(1911.11, 3))
}

func Sqrt(x float64) float64 {
	z := x

	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func Round(x float64, n int) float64 {

	x = math.Pow(10, -float64(n)) * x
	mod := math.Mod(10*x, 10)
	//fmt.Printf("%f %d %f %f ", x, n, y, mod)
	if mod <= 4 {
		return math.Floor(x) * math.Pow(10, float64(n))
	}
	return math.Ceil(x) * math.Pow(10, float64(n))
}
