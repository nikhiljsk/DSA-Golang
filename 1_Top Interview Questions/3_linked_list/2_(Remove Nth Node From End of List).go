func removeNthFromEnd(head *ListNode, n int) *ListNode {
	endPointer := head
	correctPointer := head
	for n > 0 && endPointer != nil {
		n--
		endPointer = endPointer.Next
	}
	if endPointer == nil {
		return correctPointer.Next
	}

	for endPointer.Next != nil {
		endPointer = endPointer.Next
		correctPointer = correctPointer.Next
	}
	correctPointer.Next = correctPointer.Next.Next
	return head
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 3M
// Time Implement - 15M