/*
Binary Tree Level Order Traversal
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []


Constraints:

The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
*/

// Approach one using two queues
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue, queue2 []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 || len(queue2) != 0 {
		var temp []int
		for len(queue) != 0 {
			node := queue[0]  // Front
			queue = queue[1:] // Pop
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue2 = append(queue2, node.Left)
			}
			if node.Right != nil {
				queue2 = append(queue2, node.Right)
			}
		}
		res = append(res, temp)
		for len(queue2) != 0 {
			node := queue2[0]   // Top
			queue2 = queue2[1:] // Pop
			queue = append(queue, node)
		}
	}
	return res
}

// Approach two would be use one queue, but iterate only as many times as the length of the queue
// This keeps track of the number of elements in the queue
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		queueLen := len(queue)
		var temp []int
		for i := 0; i < queueLen; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

/*
Testcases
[3,9,20,null,null,15,7]
[0,2,4,1,null,3,-1,5,1,null,6,null,8]
[1,2,3,null,4,6,7,5,null,null,null,null,8,null,10,null,9]
[1]
[]
*/

// Approach two, had to see that queueLen is used. Implemented
// Implemented