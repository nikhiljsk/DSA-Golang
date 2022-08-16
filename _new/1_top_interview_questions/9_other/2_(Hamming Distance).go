func hammingWeight(num uint32) int {
	var res int
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

func hammingDistance(x int, y int) int {
	return hammingWeight(uint32(x) ^ uint32(y))
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 1M
// Time Implement - 2M