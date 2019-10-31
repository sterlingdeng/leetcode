package apple_stocks

import "math"

/**
 *  I have an array stock_prices_yesterday where:
 *
 *    - The indices are the time in minutes past trade opening time, which was 9:30am local time
 *    - The values are the prices in dollars of Apple stock at the time
 *
 *  For example, the stock cost $500 at 10:30am, so stock_prices_yesterday[60] = 500;
 *
 *  Write an efficient algorithm for computing the best profit I could have made from 1 purchase
 *  and 1 sale of 1 Apple stock yesterday
 *
 *  Return 0 if no profit is possible OR if input is invalid.
 */

// [5, 10, 6, 20, 3, 4, 21]

func bestProfit(prices []int) int {
	best := math.Inf(-1) // first initialize a variable to hold the best profit. best will be type float64 because go >_>
	low := float64(prices[0])  // keep track of the lowest value that we have seen
	for i := 0; i < len(prices); i++ {
		curr := float64(prices[i]) // get the current value and convert to float64 >_>
		best = math.Max(best, curr-low) //
		if curr < low {
			low = curr
		}
	}
	return int(best) // type convert best back to int
}

