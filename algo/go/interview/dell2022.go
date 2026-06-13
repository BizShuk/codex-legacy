package i

// [Question]: Find top k count of english terms, one line one term.
// [Assumption]: maximum 10k count
func Dell2022_1(tokens []string, topK int) []string {

	// tokens := []string{
	// 	"Apple",
	// 	"Banana",
	// 	"Apple",
	// 	"Car",
	// 	"Banana",
	// 	"Apple",
	// }

	counter := make(map[string]int)
	buckets := make([][]string, 10000)

	for _, term := range tokens {
		counter[term] += 1
	}

	for term, count := range counter {
		buckets[count] = append(buckets[count], term)
	}

	result := make([]string, 0)

	for i := 9999; i >= 0; i-- {
		for j := 0; j < len(buckets[i]); j++ {
			result = append(result, buckets[i][j])
			topK -= 1
			if topK == 0 {
				return result
			}
		}
	}
	return result
}

// [Question]: Sort 3 type of balls in slice
// [Time  Complexity]: O(n)
// [Space Complexity]: O(1)
func Dell2022_2(balls *[]int) {
	lp, rp := 0, len(*balls)-1

	i := 0
	for i < len(*balls) && i <= rp {
		if (*balls)[i] == 2 {
			(*balls)[i], (*balls)[rp] = (*balls)[rp], (*balls)[i]
			rp -= 1
			continue
		}

		if (*balls)[i] == 0 {
			(*balls)[i], (*balls)[lp] = (*balls)[lp], (*balls)[i]
			lp += 1
		}

		i += 1
	}
}
