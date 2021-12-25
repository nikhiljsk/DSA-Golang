/*
203. Remove Linked List Elements

Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.



Example 1:


Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]
Example 2:

Input: head = [], val = 1
Output: []
Example 3:

Input: head = [7,7,7,7], val = 7
Output: []


Constraints:

The number of nodes in the list is in the range [0, 104].
1 <= Node.val <= 50
0 <= val <= 50
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	iter, prev := head, head
	for iter != nil {
		if iter.Val == val {
			if iter == head { // Handle head node deletion
				head = iter.Next
			} else {
				prev.Next = iter.Next // Other deletion
			}
		} else {
			prev = iter
		}
		iter = iter.Next
	}
	return head
}

// Implemented