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
	curr, temp := head, head
	// Get the length of LL
	count := 1
	for temp.Next != nil {
		count++
		temp = temp.Next
	}
	// Handle single node LL
	if count == 1 {
		return nil
	}
	// Handle first node deletion
	if count-n == 0 {
		return head.Next
	}
	// Get to the previous node of target
	for i := 0; i < count-n-1; i++ {
		curr = curr.Next
	}
	// Handle if target is a tail node
	if curr.Next.Next == nil {
		curr.Next = nil
		return head
	}
	// Handle if target is in the middle
	curr.Next = curr.Next.Next

	return head
}

// Implemented
