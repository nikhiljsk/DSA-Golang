// Invalid solution - Tried Inorder traversal
func isValidBST(root *TreeNode) bool {
	var inorder []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		// The solution fails here. Instead of iterating and pushing the left,
		// the solution should ideally push the root itself and then go for iteration
		// Consider the case 2, 1, 3. A loop arises 2 & 1 and it never leaves
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		} else {
			inorder = append(inorder, curr.Val)
			stack = stack[:len(stack)-1]
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
		}
	}

	fmt.Println("Inorder", inorder)
	for i := 0; i < len(inorder)-1; i++ {
		if inorder[i] > inorder[i+1] {
			return false
		}
	}
	return true
}

// Approach 1 - Inorder Iterative with entire Array
func isValidBST(root *TreeNode) bool {
	var res []int
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) != 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, curr.Val)
			curr = curr.Right
		}
	}

	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true

}

// Approach 2 - Inorder Iterative only one element
// The trick here is to use the prev as a Node so that comparision is possible
// Otherwise it'll fail for cases where int_min is given
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	var stack []*TreeNode
	curr := root
	for curr != nil || len(stack) != 0 {
		if curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if prev != nil && prev.Val >= curr.Val {
				return false
			} else {
				prev = curr
			}
			curr = curr.Right
		}
	}
	return true
}

//  Approach 3 - Recursive valid range traversal
func isValidBST(root *TreeNode) bool {
	var isValid func(root, low, high *TreeNode) bool
	isValid = func(root, low, high *TreeNode) bool {
		if root == nil {
			return true
		}
		val := root.Val
		if (low != nil && val <= low.Val) || (high != nil && val >= high.Val) {
			return false
		}
		return isValid(root.Left, low, root) && isValid(root.Right, root, high)
	}
	return isValid(root, nil, nil)
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No
// Solution - Yes
// Hints - No
// Time Approach - 30M
// Time Implement - 40M