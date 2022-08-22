func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 1 && n == 0 {
		return
	} else if m == 0 && n == 1 {
		nums1[0] = nums2[0]
		return
	}
	i, j := m-1, n-1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for k >= 0 && j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

// Same approach, cleaner code
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i, j := m-1, n-1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	if j >= 0 {
		copy(nums1[:j+1], nums2[:j+1])
	}
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 2M
// Time Implement - 15M