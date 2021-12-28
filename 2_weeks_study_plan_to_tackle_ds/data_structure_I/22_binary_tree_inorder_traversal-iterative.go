/*
Binary Tree Inorder Traversal
left, root, right
Given the root of a binary tree, return the inorder traversal of its nodes' values.



Example 1:


Input: root = [1,null,2,3]
Output: [1,3,2]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?
*/

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	for {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			if len(stack) == 0 {
				break
			}
			root = stack[len(stack)-1]   // top
			stack = stack[:len(stack)-1] //pop
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}

// Solution