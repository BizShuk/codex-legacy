package util

// [Pattern]: [Go Util] Abs
func Abs[K Number](x K) K {
	if x >= 0 {
		return x
	}
	return -x
}

func AbsDiff[K Number](x, y K) K {
	if x > y {
		return x - y
	}
	return y - x
}
