// Approach 1 - Recursion - Solution
func compare(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Val != b.Val {
		return false
	}
	return compare(a.Left, b.Right) && compare(a.Right, b.Left)
}

func isSymmetric(root *TreeNode) bool {
	return compare(root.Left, root.Right)
}

// Approach 2 - Iterative
func isSymmetric(root *TreeNode) bool {
	var s []*TreeNode
	s = append(s, root.Right)
	s = append(s, root.Left)
	for len(s) != 0 {
		l := s[len(s)-1]
		s = s[:len(s)-1]
		r := s[len(s)-1]
		s = s[:len(s)-1]

		if l == nil && r == nil {
			continue
		} else if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}

		s = append(s, r.Right)
		s = append(s, l.Left)
		s = append(s, r.Left)
		s = append(s, l.Right)
	}
	return true
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - Yes
// Hints - No
// Time Approach - 10M
// Time Implement - 25M