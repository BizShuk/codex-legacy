package slidingwindow

import (
	"math"
)

// [Pattern]: [Simple Sliding Window] find max profit in one tx
func MaxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxProfit := 0
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		}
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return -1
}
