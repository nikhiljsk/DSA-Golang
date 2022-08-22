/*
Symmetric Tree

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).



Example 1:


Input: root = [1,2,2,3,4,4,3]
Output: true
Example 2:


Input: root = [1,2,2,null,3,null,3]
Output: false


Constraints:

The number of nodes in the tree is in the range [1, 1000].
-100 <= Node.val <= 100
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Approach 1
// Storage - O(n)
// Time - O(n)
// This could be improved. BFS. Check at each level and then goto the next level
// Instead of storing all nodes, the storage complexity is reduced to O(Width of BST)
func isSymmetric(root *TreeNode) bool {
	var arr [][]int
	var s []*TreeNode
	arr = append(arr, []int{root.Val})
	s = append(s, root)

	for len(s) != 0 {
		var temp []int
		qLen := len(s)
		for i := 0; i < qLen; i++ {
			node := s[0]
			s = s[1:]
			if node.Left == nil {
				temp = append(temp, -101)
			} else {
				temp = append(temp, node.Left.Val)
				s = append(s, node.Left)
			}
			if node.Right == nil {
				temp = append(temp, -101)
			} else {
				temp = append(temp, node.Right.Val)
				s = append(s, node.Right)
			}
		}
		arr = append(arr, temp)
	}
	// fmt.Println(arr)

	// Check if symmetric
	for _, row := range arr {
		i := 0
		j := len(row) - 1
		for i < j {
			if row[i] != row[j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}

// Implemented