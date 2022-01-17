/*
Rotate Array

Given an array, rotate the array to the right by k steps, where k is non-negative.



Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]


Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105


Follow up:

Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
Could you do it in-place with O(1) extra space?
*/

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	var i, count int

	for count < n {
		pos := (i + k) % n
		curr := nums[pos]
		nums[pos] = nums[i]
		count++
		j := pos
		for count < n && j != i {
			pos = (j + k) % n
			curr, nums[pos] = nums[pos], curr
			count++
			j = pos
		}
		i++
	}
}

// The double loop is essential here. For odd we might not need. But just take a case where both
// n and k are 2. So double loop handles all cases

// Able to crack the approach and implemented till one loop. The other loop had to see solution