package op

// 將所有bits相加 e.g. 101010101 = 5
func HammingWeight(n int) int {
	ret := 0
	for i := 0; i < 32; i++ {
		ret += n & 1
		n = n >> 1
	}
	return ret
}
