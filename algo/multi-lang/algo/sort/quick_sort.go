package sort

func QuickSort(src []int) (int, int) {
	ct, et := quick_sort(src, 0, len(src)-1)
	return ct, et
}

func quick_sort(src []int, left, right int) (int, int) {
	ct, et := 0, 0
	key, value := left, src[left]
	i, j := left+1, right
	if i > j {
		return ct, et
	}
	for {
		for ; i < j && src[i] < value; i++ {
			key = i
		}
		for ; i < j && src[j] > value; j-- {
		}

		if i >= j {
			copy(src[left:], src[left+1:key+1])
			src[key] = value
			quick_sort(src, left, key)
			quick_sort(src, key+1, right)
			break
		}
		src[i], src[j] = src[j], src[i]

	}

	return ct, et
}
