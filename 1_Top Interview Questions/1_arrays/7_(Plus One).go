func plusOne(digits []int) []int {
	var carry, sum int
	carry = 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum = digits[i] + carry
		if sum > 9 {
			carry = sum - 9
			digits[i] = sum % 10
		} else {
			digits[i] = sum
			carry = 0
		}
	}
	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No (Had to see carry would be 1)
// Solution - No
// Hints - No
// Time Approach - 1M
// Time Implement - 15M