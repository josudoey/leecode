package code

// ref https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

// Example 1:
// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

// Example 2:
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	result := 0
	minPrice := prices[0]
	for _, price := range prices[1:] {
		minPrice = min(minPrice, price)
		profit := price - minPrice
		result = max(result, profit)
	}

	return result
}
