package Fibonacci

func Fibonacci_recursive(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	return Fibonacci_recursive(n-1) + Fibonacci_recursive(n-2)

}
