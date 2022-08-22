// Approach 1
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res <<= 1
		res = res | (num & 1)
		num >>= 1
	}
	return res
}

// Approach 2
// https://www.youtube.com/watch?v=-5z9dimxxmI
func reverseBits(num uint32) uint32 {
	num = ((num & 0xffff0000) >> 16) | ((num & 0x0000ffff) << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 20M
// Time Implement - 30M