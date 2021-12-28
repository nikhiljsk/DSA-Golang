/*
Binary Tree Postorder Traversal
Given the root of a binary tree, return the postorder traversal of its nodes' values.



Example 1:


Input: root = [1,null,2,3]
Output: [3,2,1]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]


Constraints:

The number of the nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?
*/

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var poot func(*TreeNode)
	poot = func(root *TreeNode) {
		if root == nil {
			return
		}
		poot(root.Left)
		poot(root.Right)
		res = append(res, root.Val)
	}
	poot(root)
	return res
}

// Implemented