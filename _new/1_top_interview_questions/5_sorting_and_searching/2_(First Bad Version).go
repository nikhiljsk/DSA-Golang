func firstBadVersion(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + ((r - l) / 2)
		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 1M
// Time Implement - 10M