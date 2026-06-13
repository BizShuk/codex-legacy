package op

func HappyNum(n int) bool {

	e := make(map[int]int)
	nn := 0
	for n != 1 {

		if _, ok := e[n]; ok {
			return false
		}
		e[n] = 1
		nn = 0
		for n > 0 {
			nn += (n % 10) * (n % 10)
			n /= 10
		}
		n = nn
	}
	return true
}

func HappyNum_induct(n int) bool {
	nn := 0
	for n >= 10 {
		nn = 0
		for n > 0 {
			nn += (n % 10) * (n % 10)
			n /= 10
		}
		n = nn
	}
	return (n == 1 || n == 7)
}

func HappyNum_recur(n int) bool {
	return false
}
