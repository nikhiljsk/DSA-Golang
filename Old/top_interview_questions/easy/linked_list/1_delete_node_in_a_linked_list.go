/*
Delete Node in a Linked List
Write a function to delete a node in a singly-linked list. You will not be given access to the head of the list, instead you will be given access to the node to be deleted directly.

It is guaranteed that the node to be deleted is not a tail node in the list.



Example 1:


Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.
Example 2:


Input: head = [4,5,1,9], node = 1
Output: [4,5,9]
Explanation: You are given the third node with value 1, the linked list should become 4 -> 5 -> 9 after calling your function.


Constraints:

The number of the nodes in the given list is in the range [2, 1000].
-1000 <= Node.val <= 1000
The value of each node in the list is unique.
The node to be deleted is in the list and is not a tail node
*/

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next

	// CONCEPT
	// Here we are just pointing the current node to the next of next node.
	// The next node will automatically be garbage collected by golang. So no memory leaks
	// And since the question mentions, NO Tail nodes, we don't need to check for nil conditions
}

// Did research to get a few things like memory leaks
// Had to see the solution. 