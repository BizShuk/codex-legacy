package bit

func Swap_by_go(x, y int) (int, int) {
	return y, x
}

func Swap_by_xor(x, y int) (int, int) {
	x = x ^ y
	y = x ^ y
	x = x ^ y
	return x, y
}
func Swap_by_plus(x, y int) (int, int) {
	x = x + y
	y = x - y
	x = x - y
	return x, y
}

// [Hint]: x, y could not be zero
func Swap_by_multiply(x, y int) (int, int) {
	x = x * y
	y = x / y
	x = x / y
	return x, y
}
