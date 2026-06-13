package slidingwindow

import "github.com/bizshuk/algo/util"

// [Pattern]: [Sliding window Two pointers] longest non-repeat char
func LengthOfLongestNonRepeatChar(s string) int {
	maxLength := 0
	set := make(map[byte]int)
	l, r := 0, 0
	for r = 0; r < len(s); r++ {
		for set[s[r]] > 0 {
			delete(set, s[l])
			l += 1
		}

		set[s[r]] += 1
		maxLength = util.Max(maxLength, r-l+1)
	}
	return maxLength
}
