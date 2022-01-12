/*
Invert Binary Tree
Given the root of a binary tree, invert the tree, and return its root.



Example 1:


Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
Example 2:


Input: root = [2,1,3]
Output: [2,3,1]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*/

// Approach 1 - Recursion
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	_ = invertTree(root.Left)
	_ = invertTree(root.Right)

	// Instead of the above three lines. You can use
	// root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)

	return root
}

// Approach 2 - Iterative
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var s []*TreeNode
	s = append(s, root)
	for len(s) != 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]

		if node != nil {
			node.Left, node.Right = node.Right, node.Left
			s = append(s, node.Left)
			s = append(s, node.Right)
		}
	}
	return root
}

/*
Testcases
[1,2,2,3,4,4,3]
[1,2,2,null,3,null,3]
[0,2,4,1,null,3,-1,5,1,null,6,null,8]
[3,9,20,null,null,15,7]
[0,2,4,1,null,3,-1,5,1,null,6,null,8]
[1,2,3,null,4,6,7,5,null,null,null,null,8,null,10,null,9]
[1]
[1,1,1]
[1,1,1,2,2,2,2]
[1,1,1,null,2,null,2]
[1,1,1,2,null,null,2]
[1,2,3]
[9,-42,-42,null,76,76,null,null,13,null,13]
*/

// Implemented