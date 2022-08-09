func intersect(nums1 []int, nums2 []int) []int {
	var nums []int
	find := make(map[int]int)
	// Iterate and create hashmap freq for first array
	for _, v := range nums1 {
		if _, ok := find[v]; ok {
			find[v] += 1
		} else {
			find[v] = 1
		}
	}
	// Iterate over second array and decrease the frequency
	for _, v := range nums2 {
		if _, ok := find[v]; ok && find[v] != 0 {
			nums = append(nums, v)
			find[v] -= 1
		}
	}
	return nums
}

// Other Approaches
// Naive n-square approach
// If sorted arrays, point i & j to each array start, and compare and do either i++, or j++ or append.

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 1M
// Time Implement - 5M