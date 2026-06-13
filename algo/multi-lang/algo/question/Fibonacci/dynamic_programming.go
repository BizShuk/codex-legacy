package Fibonacci

var tmp = make(map[int]int)

func Fibonacci_dynamicprogramming(n int) int {
	if n == 0 {
		tmp[n] = 0
	}

	if n == 1 {
		tmp[n] = 1
	}
	if _, ok := tmp[n]; ok {
		return tmp[n]
	}

	tmp[n] = Fibonacci_dynamicprogramming(n-2) + Fibonacci_dynamicprogramming(n-1)
	return tmp[n]
}
