func containsDuplicate(nums []int) bool {
	find := make(map[int]bool)
	for _, i := range nums {
		if _, ok := find[i]; ok {
			return true
		}
		find[i] = true
	}
	return false
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 1M
// Time Implement - 3M