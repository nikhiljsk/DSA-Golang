func removeDuplicates(nums []int) int {
	prev := 0
	for curr := 1; curr < len(nums); curr++ {
		if nums[prev] != nums[curr] {
			nums[prev+1] = nums[curr]
			prev++
		}
	}
	return prev + 1 // since index
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 1M
// Time Implement - 6M