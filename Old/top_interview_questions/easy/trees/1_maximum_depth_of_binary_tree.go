/*
Maximum depth of binary tree
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 3
Example 2:

Input: root = [1,null,2]
Output: 2


Constraints:

The number of nodes in the tree is in the range [0, 104].
-100 <= Node.val <= 100
*/

func maxDepth(root *TreeNode) int {
	var count int
	if root == nil {
		return count
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		qlen := len(queue)
		count++
		for i := 0; i < qlen; i++ {
			node := queue[0]  // Front
			queue = queue[1:] // Pop
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return count
}

// Approach 2 using recursion
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Right), maxDepth(root.Left))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Testcases
[3,9,20,null,null,15,7]
[0,2,4,1,null,3,-1,5,1,null,6,null,8]
[1,2,3,null,4,6,7,5,null,null,null,null,8,null,10,null,9]
[1]
[]
*/

// Approach 2 Solution
// Implemented