package binerysearch

import "github.com/bizshuk/algo/util"

// [Variant]: [binary search] slowest eating speed
// Usually, binary search will stop if target hit. This case won't stop if condition matched, it must find mininum speed under target hour
// [Hint]: match and moving condition, especially for equal scenario.
// When required hour matched( <= h), slow speed (rp = mp -1) to find more smaller solutions
func MinEatingSpeed(piles []int, h int) int {
	minEatingRate := 0

	for _, pile := range piles {
		minEatingRate = util.Max(minEatingRate, pile)
	}

	// Find min of eating rate
	// ex3: 23 is min, but 24 25 26 27 28 29 is finished withing 6 hour
	lp, rp := 1, minEatingRate

	for lp <= rp {
		speed := (lp + rp) / 2
		hour := EatingTime(piles, speed)

		// h = hour => slow speed => rp = speed -1
		if h >= hour {
			minEatingRate = util.Min(minEatingRate, speed)
			rp = speed - 1
		} else {
			lp = speed + 1
		}

	}

	return minEatingRate
}

func EatingTime(piles []int, speed int) (result int) {
	for _, pile := range piles {
		result += pile / speed
		if pile%speed != 0 {
			result += 1
		}
	}
	return result
}
