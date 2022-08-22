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

// Approach 1 - Using two stacks
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var s1, s2 []*TreeNode
	s1 = append(s1, root)
	for len(s1) != 0 {
		node := s1[len(s1)-1] // top
		s1 = s1[:len(s1)-1]   // pop
		s2 = append(s2, node)
		if node.Left != nil {
			s1 = append(s1, node.Left)
		}
		if node.Right != nil {
			s1 = append(s1, node.Right)
		}
	}
	for len(s2) != 0 {
		node := s2[len(s2)-1] // top
		s2 = s2[:len(s2)-1]   // pop
		res = append(res, node.Val)
	}
	return res
}

// Approach 2 Using 1 Stack
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) != 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			temp := stack[len(stack)-1].Right
			if temp != nil {
				curr = temp
			} else {
				temp = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res = append(res, temp.Val)
				for len(stack) != 0 && temp == stack[len(stack)-1].Right {
					temp = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					res = append(res, temp.Val)
				}
			}
		}
	}
	return res
}

// Solution