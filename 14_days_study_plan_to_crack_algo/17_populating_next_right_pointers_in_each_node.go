/*
Populating Next Right Pointers in Each Node
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.



Example 1:


Input: root = [1,2,3,4,5,6,7]
Output: [1,#,2,3,#,4,5,6,7,#]
Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
Example 2:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 212 - 1].
-1000 <= Node.val <= 1000


Follow-up:

You may only use constant extra space.
The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.
*/

// Approach 1 - BFS
// Time complexity: O(N)
// Space complexity: O(N/2) - Width
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var q []*Node
	q = append(q, root)

	for len(q) != 0 {
		qLen := len(q) // Record length for level-order traversal

		prev := q[0] // Get front
		q = q[1:]

		if prev == nil { // If nil found, full tree is processed (Since complete binary tree)
			break
		}

		q = append(q, prev.Left)
		q = append(q, prev.Right)

		for i := 1; i < qLen; i++ {
			curr := q[0]
			q = q[1:]

			prev.Next = curr
			prev = curr

			q = append(q, curr.Left)
			q = append(q, curr.Right)
		}
		prev.Next = nil
	}
	return root
}

// Approach 2
func connect(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	// Deal with nodes having common parent node
	for curr := root; curr != nil; curr = curr.Next {
		curr.Left.Next = curr.Right
	}

	// Deal with nodes NOT having common parent node
	for curr := root; curr != nil && curr.Next != nil; curr = curr.Next {
		curr.Right.Next = curr.Next.Left
	}

	connect(root.Left)
	return root

}

// Implemented approach 1
// Solution for approach 2