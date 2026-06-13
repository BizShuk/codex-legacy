package Fibonacci

func Fibonacci_iterative(n int) int {

	var v1, v2 int = 0, 1
	var result = 0
	for i := 0; i < n-1; i++ {
		if i == 0 {
			result = v1
		}
		if i == 1 {
			result = v2
		}

		result = v1 + v2
		v1 = v2
		v2 = result
	}

	return result
}
