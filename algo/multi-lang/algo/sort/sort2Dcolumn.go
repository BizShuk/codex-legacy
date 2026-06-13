package sort

func Sort2Dcolumn(s [][]int) [][]int {

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			a := s[i]
			b := s[j]
			x := 0
			first_longer := len(a) > len(b)

			for {
				if len(a) <= x || len(b) <= x {
					if first_longer {
						s[i], s[j] = s[j], s[i]
					}
					break
				}

				if a[x] > b[x] {
					s[i], s[j] = s[j], s[i]
					break
				}

				if a[x] < b[x] {
					break
				}
				x++
			}

		}
	}
	return s
}
