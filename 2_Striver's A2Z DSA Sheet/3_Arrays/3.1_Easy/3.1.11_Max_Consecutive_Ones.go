func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxConsecutiveOnes(nums []int) int {
	global, local := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			local++
			global = max(global, local)
		} else {
			local = 0
		}
	}
	return global
}