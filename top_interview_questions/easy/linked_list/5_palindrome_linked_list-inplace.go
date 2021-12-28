/*
Palindrome Linked List

Given the head of a singly linked list, return true if it is a palindrome.



Example 1:


Input: head = [1,2,2,1]
Output: true
Example 2:


Input: head = [1,2]
Output: false


Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9


Follow up: Could you do it in O(n) time and O(1) space?
*/

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

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	// Iterate m to middle of linked list
	l, m, r := head, head, head
	for r.Next != nil && m.Next != nil && r.Next.Next != nil {
		r = r.Next.Next
		m = m.Next
	}

	rev_m := reverseList(m.Next) // Reverse the linked list from m+1
	m.Next = rev_m               // Do inplace for LL
	m = m.Next                   // Now m = m+1 for later comparision

	// Check if l and m values are equal continuously
	for m != nil && l != nil {
		if l.Val != m.Val {
			return false
		}
		m = m.Next
		l = l.Next
	}

	return true

}

// Half approach cracked. Solution for later half. Implemented on my own