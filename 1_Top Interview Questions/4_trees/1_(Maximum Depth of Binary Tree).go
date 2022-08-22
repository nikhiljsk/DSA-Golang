// Approach 1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var q []*TreeNode
	q = append(q, root)
	var depth int
	for len(q) != 0 {
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			curr := q[0]
			q = q[1:]
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		depth++
	}
	return depth
}

// Approach 2 - Solution
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No (qLen)
// Solution - Yes
// Hints - No
// Time Approach - 15M
// Time Implement - 35M