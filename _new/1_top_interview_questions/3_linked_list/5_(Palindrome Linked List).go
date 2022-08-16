func reverseLL(head *ListNode) *ListNode {
	if head == nil {
		return nil
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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return head != nil
	}
	middle, end := head, head
	for end.Next != nil && end.Next.Next != nil {
		end = end.Next.Next
		middle = middle.Next
	}
	middle.Next = reverseLL(middle.Next)
	middle = middle.Next
	for head != nil && middle != nil {
		if head.Val != middle.Val {
			return false
		}
		head = head.Next
		middle = middle.Next
	}
	return true
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 1M
// Time Implement - 15M