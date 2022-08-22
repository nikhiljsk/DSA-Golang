func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var q []*TreeNode
	var res [][]int
	q = append(q, root)

	for len(q) != 0 {
		var temp []int
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			curr := q[0]
			q = q[1:]
			temp = append(temp, curr.Val)
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

// Approach 2 - Solution
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var helper func(root *TreeNode, depth int)

	helper = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) <= depth {
			res = append(res, []int{root.Val})
		} else {
			res[depth] = append(res[depth], root.Val)
		}
		helper(root.Left, depth+1)
		helper(root.Right, depth+1)
	}

	helper(root, 0)
	return res
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - Yes
// Hints - No
// Time Approach - 10M
// Time Implement - 25M