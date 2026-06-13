package array

import "sort"

// [Pattern]: [Go Sorting] string swap, sort.Slice, hash slice
// 1. string CAN'T swap in-place
// 2. sort.Slice(sliceVar, func(i int, j int) bool) is in-place
// 3. string([]rune | []byte) are hash kind value

func FindGroupAnagrams(strs []string) [][]string {
	anagrams := make([][]string, 0)
	charCounter := make(map[string][]string)

	// [Solution] 1: sort string and match
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i int, j int) bool {
			return bytes[i] < bytes[j]
		})
		sortStr := string(bytes)
		charCounter[sortStr] = append(charCounter[sortStr], str)
	}

	// [Solution] 2: sort alphabetic rune/byte slice and compare slices

	// [Solution] 3: hash rune/byte slice by string([]rune/[]byte)
	// for _, str := range strs {
	//     hash := hashWord(str)
	//     charCounter[hash] = append(charCounter[hash], str)
	// }

	for _, v := range charCounter {
		anagrams = append(anagrams, v)
	}

	return anagrams
}

// [Pattern]: [Hash by rune/byte slice] string([]rune), rune/byte count limit is int32/int8
func HashWord(str string) string {
	x := make([]rune, 26)
	for _, c := range str {
		x[c-97] += 1
	}
	return string(x)
}
