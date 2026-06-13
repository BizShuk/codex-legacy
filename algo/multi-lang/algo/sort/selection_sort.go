package sort

// select max(min) value in this round and put it in the head(tail)
func SelectionSort(src []int) (int, int) {
	var ct, et, tmp int
	for i := 0; i < len(src); i++ {
		tmp = i
		for j := i + 1; j < len(src); j++ {
			ct++
			if src[tmp] > src[j] {
				tmp = j
			}
		}
		if tmp != i {
			src[i], src[tmp] = src[tmp], src[i]
			et++
		}
	}
	return ct, et
}
