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

// Approach 2
// Storage - O(h)
// Time - O(n)
func isSymmetric(root *TreeNode) bool {
	var s1, s2 []*TreeNode
	s2 = append(s2, root.Right)
	s1 = append(s1, root.Left)

	for len(s1) != 0 && len(s2) != 0 {
		lst := s1[len(s1)-1]
		rst := s2[len(s2)-1]
		s1 = s1[:len(s1)-1]
		s2 = s2[:len(s2)-1]

		if lst == nil && rst == nil {
			continue
		} else if lst == nil || rst == nil {
			return false
		}

		if lst.Val != rst.Val {
			return false
		}

		s1 = append(s1, lst.Left)
		s2 = append(s2, rst.Right)
		s1 = append(s1, lst.Right)
		s2 = append(s2, rst.Left)

	}
	return len(s1) == 0 && len(s2) == 0
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