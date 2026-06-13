package sort

// pick next and insert to sorted slice
// important: how to get j which is point of stopping comparsion
func InsertionSort(src []int) (int, int) {
	ct, et := 0, 0
	var tmp int

	for i := 1; i < len(src); i++ {
		tmp = src[i]

		for j := i - 1; ; j-- {
			ct++
			if j == -1 || tmp > src[j] {
				copy(src[j+2:], src[j+1:i])
				src[j+1] = tmp
				et++
				break
			}
		}
	}
	return ct, et
}
