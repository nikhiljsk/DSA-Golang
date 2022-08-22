func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	superFast := head.Next
	for head != nil && superFast != nil && superFast.Next != nil {
		superFast = superFast.Next.Next
		head = head.Next
		if head == superFast {
			return true
		}
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
// Time Implement - 2M