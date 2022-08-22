/*
Remove Duplicates from Sorted List

Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.



Example 1:


Input: head = [1,1,2]
Output: [1,2]
Example 2:


Input: head = [1,1,2,3,3]
Output: [1,2,3]


Constraints:

The number of nodes in the list is in the range [0, 300].
-100 <= Node.val <= 100
The list is guaranteed to be sorted in ascending order.
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	seen := make(map[int]bool)
	prev, iter := head, head
	for iter != nil {
		if _, ok := seen[iter.Val]; ok {
			// Handle deletion of node
			if iter == head { // First node deletion
				head = iter.Next
			} else {
				prev.Next = iter.Next
			}
		} else {
			seen[iter.Val] = true
			prev = iter
		}
		iter = iter.Next
	}
	return head
}

// Approach two - O(1) Storage
// Use a lastSeen value variable instead of hashmap as it is already sorted
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lastSeen := -101
	prev, iter := head, head
	for iter != nil {
		if lastSeen == iter.Val {
			// Handle deletion of node
			if iter == head { // First node deletion
				head = iter.Next
			} else {
				prev.Next = iter.Next
			}
		} else {
			lastSeen = iter.Val
			prev = iter
		}
		iter = iter.Next
	}
	return head
}

/*
Testcases
[1,1,2]
[1,1,2,3,3]
[]
[1]
[1,1,1,1]
[1,2,3,4,4,4,4,5,5]
[1,1]
*/

// Implemented both