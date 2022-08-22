func twoSum(nums []int, target int) []int {
	find := make(map[int]int)
	for i, v := range nums {
		if j, ok := find[v]; ok && i != j {
			return []int{j, i}
		}
		find[target-v] = i
	}
	return []int{}
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 10M (Got confused that the array is sorted or can be sorted)
// Time Implement - 10M