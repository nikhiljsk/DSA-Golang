func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var l *ListNode
	m, r := head, head.Next
	for r != nil {
		m.Next = l
		l = m
		m = r
		r = r.Next
	}
	m.Next = l
	return m
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No
// Solution - No
// Hints - No
// Time Approach - 3M
// Time Implement - 5M