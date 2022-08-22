/*
121. Best Time to Buy and Sell Stock

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.



Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.


Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104
*/

// Approach 1 - Faster than approach two
// Go through index and index+1, if index < index+1 it's not worth the buy continue, otherwise basket index
// Keep refreshing the basked with lowestvalue and get maxProfit of current index
func maxProfit(prices []int) int {
	buy := 100000
	var maxProfit int

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] && buy > prices[i] {
			buy = prices[i]
			continue
		}
		if maxProfit < prices[i]-buy {
			maxProfit = prices[i] - buy
		}
	}
	if maxProfit < prices[len(prices)-1]-buy {
		maxProfit = prices[len(prices)-1] - buy
	}
	return maxProfit
}

// Approach 2 - Go through each index, and then basket the one which is lower than the previous index.
// For each index, calculate and take the max profit with the basketed value
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	for _, p := range prices {
		tmp := p - buy
		if tmp <= 0 {
			buy = p
		}
		profit = max(p-buy, profit)
	}
	return profit
}

// Able to crack the approach and implemented