package main

import "math"

func maxProfit(prices []int) int {

	l, r := 0, 1
	maxProf := 0

	for r < len(prices) {
		if prices[l] > prices[r] {
			l = r
		} else {
			profit := prices[r] - prices[l]
			maxProf = int(math.Max(float64(maxProf), float64(profit)))
		}

		r++
	}

	return maxProf

}
