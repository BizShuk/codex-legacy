package op

func AddDigits(n int) int {
	var ret int

	for n >= 10 {
		ret += n % 10
		n = n / 10
	}
	ret += n % 10
	if ret >= 10 {
		return AddDigits(ret)
	}
	return ret
}

func AddDigits_math(n int) int {
	return 1 + (n-1)%9
}
