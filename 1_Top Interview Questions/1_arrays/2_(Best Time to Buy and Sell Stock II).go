func maxProfit(prices []int) int {
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += (prices[i+1] - prices[i])
		}
	}
	return profit
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 1M
// Time Implement - 4M