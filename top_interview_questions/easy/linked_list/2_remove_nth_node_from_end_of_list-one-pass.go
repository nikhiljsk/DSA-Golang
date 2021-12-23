/*
Remove Nth Node From End of List

Given the head of a linked list, remove the nth node from the end of the list and return its head.



Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]


Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


Follow up: Could you do this in one pass?
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr, temp, i := head, head, 0
	// Iterate, till n is reached
	for i < n-1 && temp.Next != nil {
		temp = temp.Next
		i++
	}
	// Handle if target is head node
	if temp == head && temp.Next == nil {
		return head.Next
	}
	// Handle if target is tail node
	if temp == curr {
		for curr.Next.Next != nil {
			curr = curr.Next
		}
		curr.Next = nil
		return head
	}
	// Move both pointers together
	for temp.Next != nil {
		curr = curr.Next
		temp = temp.Next
	}
	// Handle if target is in the middle
	curr.Val = curr.Next.Val
	curr.Next = curr.Next.Next
	return head
}

// Approach 2
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	preHead := &ListNode{Next: head}
	l, r := preHead, head
	for i := 0; i < n && r != nil; i++ {
		r = r.Next
	}
	for r != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next
	return preHead.Next
}

// Implemented approach 1
// Solution for approach 2 using preheader