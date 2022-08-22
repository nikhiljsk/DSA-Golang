// Approach 1
func deleteNode(node *ListNode) {
	head := node
	for head.Next.Next != nil {
		head.Val = head.Next.Val
		head = head.Next
	}
	head.Val = head.Next.Val
	head.Next = nil
}

// Approach 2
// Here we are just pointing the current node to the next of next node.
// The next node will automatically be garbage collected by golang. So no memory leaks
// And since the question mentions, NO Tail nodes, we don't need to check for nil conditions
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 3M
// Time Implement - 10M