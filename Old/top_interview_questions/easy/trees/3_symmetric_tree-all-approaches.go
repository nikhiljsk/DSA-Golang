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

// Approach 3 - Recursion
// Storage - O(h)
// Time - O(n)
func isTrue(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isTrue(a.Left, b.Right) && isTrue(a.Right, b.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isTrue(root.Left, root.Right)
}

// Approach 4 - Iterative - One Stack
// Storage - O(h)
// Time - O(n)
func isSymmetric(root *TreeNode) bool {
	var s []*TreeNode
	// Push left and right
	s = append(s, root.Right)
	s = append(s, root.Left)

	for len(s) != 0 {
		// Pop values out of stack
		lst := s[len(s)-1]
		s = s[:len(s)-1]
		rst := s[len(s)-1]
		s = s[:len(s)-1]

		// Check for nil
		if lst == nil && rst == nil {
			continue
		} else if lst == nil || rst == nil {
			return false
		}

		// Check for value
		if lst.Val != rst.Val {
			return false
		}

		// Push into stack so that, mirror condition is satisfied
		// lst.Left == rst.Right && lst.Right == rst.Left
		s = append(s, rst.Right)
		s = append(s, lst.Left)
		s = append(s, rst.Left)
		s = append(s, lst.Right)
	}
	return true
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

// Approach 1, 2, 4, Implemented
// Solution for approach 3 - Recursion
