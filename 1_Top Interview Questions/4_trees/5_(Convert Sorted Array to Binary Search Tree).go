// Invalid Approach - Was trying for a cone like structure :|
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0], Left: nil, Right: nil}
	} else if len(nums) == 2 {
		l := TreeNode{Val: nums[0], Left: nil, Right: nil}
		return &TreeNode{Val: nums[1], Left: &l, Right: nil}
	}
	mid := len(nums) / 2
	rootVal := TreeNode{Val: nums[mid], Left: nil, Right: nil}
	root := &rootVal
	leftIter, rightIter := root, root
	for i := 1; mid+i < len(nums); i++ {
		lv, rv := TreeNode{}, TreeNode{}
		l, r := &lv, &rv
		l.Val = nums[mid-i]
		r.Val = nums[mid+i]
		leftIter.Left = l
		rightIter.Right = r
		leftIter = l
		rightIter = r
	}
	if mid%2 == 0 {
		tv := TreeNode{}
		t := &tv
		t.Val = nums[0]
		leftIter.Left = t
	}
	return root
}

// Approach 1 - Solution
func buildBST(left, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + ((right - left) / 2)
	node := &TreeNode{Val: nums[mid]}
	node.Left = buildBST(left, mid-1, nums)
	node.Right = buildBST(mid+1, right, nums)
	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(0, len(nums)-1, nums)
}

// Same approach as above, the Golang way
func buildBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  buildBST(nums[:mid]),
		Right: buildBST(nums[mid+1:]),
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums)
}

// ** Metadata **
// Intuition - No
// Approach - No
// Implementation - No
// Solution - Yes
// Hints - No
// Time Approach - 40M
// Time Implement - 40M