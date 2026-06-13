package op

func UglyNum(n int) bool {
	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}

	if n == 1 {
		return true
	}
	return false
}

// 2倍數 bit規律 尾數0
// 3倍數 bit規律 ??
// 5倍數 bit規律 ??
