package sort

// kpi: compare times , exchange times , times,

// exchange each two of value , at least one is sorted in one round
func BubbleSort(src []int) (int, int) {
	ct, et := 0, 0
	for i := 0; i < len(src); i++ {
		for j := i + 1; j < len(src); j++ {
			ct++
			if src[i] > src[j] {
				src[i], src[j] = src[j], src[i]
				et++
			}
		}
	}
	return ct, et
}
