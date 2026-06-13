package dy

var case8 = []int{1, 2, 3, 4, 5}

// multiple transactions
func MaxProfit_multi_transactions(prices []int) int {
	var max, stock_price int
	var haveone bool = false
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] >= prices[i+1] {
			if haveone {
				max += prices[i] - stock_price
				stock_price = 0
				haveone = false
			}
		} else {
			if !haveone {
				stock_price = prices[i]
				haveone = true
			}
		}
	}
	if haveone {
		max += prices[len(prices)-1] - stock_price
	}
	return max
}
