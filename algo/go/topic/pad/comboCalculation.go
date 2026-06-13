package pad

func ComboCalculate(matrix [][]int, status map[int]int) map[int]int {
	var summary map[int]int
	for {
		counter, current_summary := EliminateCombo(matrix)
		summary = CombineSummary(summary, current_summary)
		Failling(matrix)
		// Filling(matrix)

		if counter == 0 {
			break
		}

		if status[ZERODROP] > 0 {
			break
		}
	}
	return summary
}

func Failling(matrix [][]int) {
	rSize, cSize := len(matrix), len(matrix[0])
	for i := 0; i < cSize; i++ {
		lp, rp := 0, 0

		for rp < rSize {
			if matrix[i][rp] != BEAD_EMPTY {
				matrix[i][lp] = matrix[i][rp]
				lp += 1
			}
			rp += 1
		}

		for lp < rSize {
			matrix[i][lp] = BEAD_EMPTY
			lp += 1
		}
	}
}

func Filling(matrix [][]int) {

}

func EliminateCombo(matrix [][]int) (int, map[int]int) {
	counter, summary := 0, map[int]int{}

	return counter, summary
}

func CombineSummary(s1 map[int]int, s2 map[int]int) map[int]int {
	for k, v := range s2 {
		s1[k] += v
	}
	return s1
}
