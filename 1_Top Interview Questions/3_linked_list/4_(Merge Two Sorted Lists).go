// Approach 1
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head1, head2 := list1, list2
	var newHead, newTail *ListNode
	if head1.Val < head2.Val {
		newHead = head1
		newTail = head1
		head1 = head1.Next
	} else {
		newHead = head2
		newTail = head2
		head2 = head2.Next
	}

	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			newTail.Next = head1
			head1 = head1.Next
		} else if head1.Val > head2.Val {
			newTail.Next = head2
			head2 = head2.Next
		}
		newTail = newTail.Next
	}

	if head1 == nil {
		newTail.Next = head2
	} else {
		newTail.Next = head1
	}
	return newHead
}

// Approach 2
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := ListNode{}
	newTail := &preHead

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			newTail.Next = list1
			list1 = list1.Next
		} else {
			newTail.Next = list2
			list2 = list2.Next
		}
		newTail = newTail.Next
	}

	if list1 == nil {
		newTail.Next = list2
	} else {
		newTail.Next = list1
	}
	return preHead.Next
}

// Approach 2 is cleaner, but needs a intuition to get the preheader concept.

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 2M
// Time Implement - 25M