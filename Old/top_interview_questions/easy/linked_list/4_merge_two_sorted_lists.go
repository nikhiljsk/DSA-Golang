/*
21. Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/

// Approach one, create a new list. O(n+m) storage
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Create a new head for new list
	one, two := list1, list2
	newHead, newTail := &ListNode{}, &ListNode{}
	if one.Val <= two.Val {
		newHead.Val = one.Val
		one = one.Next
	} else {
		newHead.Val = two.Val
		two = two.Next
	}
	newTail = newHead

	// Iterate over both lists and assign accordingly
	for one != nil && two != nil {
		temp := &ListNode{}
		if one.Val <= two.Val {
			temp.Val = one.Val
			one = one.Next
		} else {
			temp.Val = two.Val
			two = two.Next
		}
		temp.Next = nil
		newTail.Next = temp
		newTail = temp
	}

	// Assign remaining nodes after comparision
	for one != nil {
		temp := &ListNode{}
		temp.Val = one.Val
		temp.Next = nil
		newTail.Next = temp
		newTail = temp
		one = one.Next
	}
	// Assign remaining nodes after comparision
	for two != nil {
		temp := &ListNode{}
		temp.Val = two.Val
		temp.Next = nil
		newTail.Next = temp
		newTail = temp
		two = two.Next
	}

	return newHead
}

// Approach 2 - Inplace - Splicing - Interweaving
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	one, two := list1, list2
	newHead, newTail := &ListNode{}, &ListNode{}

	// Create a new head for new list
	if one.Val <= two.Val {
		newHead = one
		one = one.Next
	} else {
		newHead = two
		two = two.Next
	}
	newTail = newHead

	// Iterate over both lists and assign accordingly
	for one != nil || two != nil {
		if one == nil { // Remaining nodes, assign all in one go and break
			newTail.Next = two
			break
		} else if two == nil { // Remaining nodes, assign all in one go and break
			newTail.Next = one
			break
		}
		if one.Val <= two.Val {
			newTail.Next = one
			newTail = one
			one = one.Next
		} else {
			newTail.Next = two
			newTail = two
			two = two.Next
		}
	}
	return newHead
}

// Implemented two approaches