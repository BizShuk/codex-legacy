package fundation

// [Pattern]: [Go Map] Created/Removed key during for-range
// [Hint]: If key is removed from map in for-range, the entry won't be processed.
// [Hint]: If key is created to map in for-range, the entry may be processed or skipped.
func MapIncrementInLoop() {
	y := map[int]bool{0: true}
	// y := make(map[int]bool)
	// y[0] = true
	i := 0
	for k := range y {
		i += 1
		y[k+1] = true
		if i > 10 {
			break
		}
	}
}
